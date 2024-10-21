package cloudflare

import (
	"bytes"
	"cf-ddns/bu_type"
	"cf-ddns/util/address_util"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type cloudflareDDNSClient struct {
	dnsName string
	zoneID  string
	token   string

	dnsID        string
	configuredIP string
	currentIP    string
}

func NewCFClient(dnsName, zoneID, token string) *cloudflareDDNSClient {
	return &cloudflareDDNSClient{
		dnsName: dnsName,
		zoneID:  zoneID,
		token:   token,
	}
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
		_this.zoneID, _this.dnsName,
	)
	req, err := http.NewRequest("GET", getRecordURL, nil)
	if err != nil {
		return
	}

	// 为请求添加自定义的头部信息
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", _this.token))
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
		return errors.New(fmt.Sprintf("no record found for %v", _this.dnsName))
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

	jsonData, err := json.Marshal(bu_type.H{
		"type":    dnsType,
		"name":    _this.dnsName,
		"content": _this.currentIP,
		"proxied": proxied,
	})
	if err != nil {
		return
	}
	req, err := http.NewRequest("PUT",
		fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%v/dns_records/%v", _this.zoneID, _this.dnsID),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return
	}

	// 为请求添加自定义的头部信息
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", _this.token))
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
