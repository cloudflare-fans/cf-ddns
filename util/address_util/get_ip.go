package address_util

import (
	"cf-ddns/bu_const"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
)

func GetIP() (ip string, err error) {
	// 发送 HTTP GET 请求
	resp, err := http.Get("https://icanhazip.com")
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

	ip = string(body)

	return
}

func GetIPDNSType(ipStr string) (dnsType bu_const.DNSType, err error) {
	log.Printf("ipStr: %v", ipStr)
	ip := net.ParseIP(ipStr)
	log.Printf("ip: %v", ip)
	if ip == nil {
		return bu_const.DNSTypeInvalid, errors.New("invalid ip")
	} else if ipv4 := ip.To4(); ipv4 != nil {
		return bu_const.DNSTypeIPv4, nil
	} else {
		return bu_const.DNSTypeIPv6, nil
	}
}
