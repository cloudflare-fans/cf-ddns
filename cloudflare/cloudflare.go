package cloudflare

import (
	"bytes"
	"cf-ddns/bu_const"
	"cf-ddns/bu_type"
	"cf-ddns/util/address_util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
)

type cloudflareDDNSClient struct {
	DNSName         string `json:"dns_name,omitempty" yaml:"dns_name,omitempty"`
	ZoneID          string `json:"zone_id,omitempty" yaml:"zone_id,omitempty"`
	Token           string `json:"token,omitempty" yaml:"token,omitempty"`
	UpdateEvery     string `json:"update_every,omitempty" yaml:"update_every,omitempty"`
	IPDetectionRule struct {
		IPType bu_const.IPType `json:"ip_type,omitempty" yaml:"ip_type,omitempty"`
	} `json:"ip_detection_rules,omitempty" yaml:"ip_detection_rules,omitempty"`

	dnsID        string
	configuredIP string
	currentIP    string
}

type config struct {
	initialized bool
	Targets     []cloudflareDDNSClient `json:"targets,omitempty" yaml:"targets,omitempty"`
}

var GlobalConfig config

func (_this *config) readConfig(configFilePath string) error {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, _this)
	if err != nil {
		return err
	}
	return nil
}

func (_this *config) InitConfig(configFilePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Config file modified or created, reloading...")
					err := _this.readConfig(configFilePath)
					if err != nil {
						log.Fatal(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = _this.readConfig(configFilePath)
	if err != nil {
		log.Fatal("Error reading initial config:", err)
	}
	log.Println("Initial config loaded successfully")
	_this.initialized = true
}

// detectCurrentIP 通过云端接口探测现在的 IP 地址
func (_this *cloudflareDDNSClient) detectCurrentIP() (err error) {
	ip, err := address_util.GetIP()
	if err != nil {
		return
	}
	_this.currentIP = ip
	return
}

// detectConfiguredRecord 获取当前的 DNS 记录 id，当前名称的 DNS 记录内容（即旧 IP）
func (_this *cloudflareDDNSClient) detectConfiguredRecord() (err error) {
	getRecordURL := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%v/dns_records?name=%v",
		_this.ZoneID, _this.DNSName,
	)
	req, err := http.NewRequest("GET", getRecordURL, nil)
	if err != nil {
		return
	}

	// 为请求添加自定义的头部信息
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", _this.Token))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Accept", "application/json")

	// 创建一个 HTTP 客户端来发送请求
	client := &http.Client{}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 处理响应体
	respBody := bu_type.GetDNSRecordRespBody{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return
	}
	if !respBody.Success {
		err = errors.New("response not success")
		return
	}

	if len(respBody.Result) == 0 {
		return errors.New(fmt.Sprintf("no record found for %v", _this.DNSName))
	}

	if len(respBody.Result) > 0 {
		firstResult := respBody.Result[0]
		_this.dnsID = firstResult.Id
		_this.configuredIP = firstResult.Content
	}
	return
}

// updateDNSRecord 更新 DNS 记录
func (_this *cloudflareDDNSClient) updateDNSRecord(proxied bool) (err error) {
	dnsType, err := address_util.GetIPDNSType(_this.currentIP)
	if err != nil {
		log.Printf("err: %v\n", err.Error())
		return
	}

	if dnsType == bu_const.DNSTypeIPv6 && _this.IPDetectionRule.IPType == bu_const.IPTypeIPv4Only {
		return errors.New("ipv4-only mode does not support detected IPv6 address")
	}

	if dnsType == bu_const.DNSTypeIPv4 && _this.IPDetectionRule.IPType == bu_const.IPTypeIPv6Only {
		return errors.New("ipv6-only mode does not support detected IPv4 address")
	}

	jsonData, err := json.Marshal(bu_type.H{
		"type":    dnsType,
		"name":    _this.DNSName,
		"content": _this.currentIP,
		"proxied": proxied,
	})
	if err != nil {
		return
	}
	req, err := http.NewRequest("PUT",
		fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%v/dns_records/%v", _this.ZoneID, _this.dnsID),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return
	}

	// 为请求添加自定义的头部信息
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", _this.Token))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Accept", "application/json")

	// 创建一个 HTTP 客户端来发送请求
	client := &http.Client{}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	respBody := bu_type.PutDNSRecordRespBody{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return
	}

	if !respBody.Success {
		err = errors.New("response not success")
		return
	}

	return
}

// shouldUpdate 是否应该更新
func (_this *cloudflareDDNSClient) shouldUpdate() bool {
	log.Print("currentIP:")
	log.Println(_this.currentIP)
	log.Print("configuredIP:")
	log.Println(_this.configuredIP)

	return _this.currentIP != "" &&
		_this.configuredIP != "" &&
		_this.currentIP != _this.configuredIP
}

// WriteLog 写入记录
//func (_this *cloudflareDDNSClient) WriteLog() (err error) {
//
//}
