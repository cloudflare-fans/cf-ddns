package main

import (
	"cf-ddns/cloudflare"
	"cf-ddns/config"
	"fmt"
)

var configFilePath = "config.yaml"

func init() {
	config.InitConfig(configFilePath)
}

func main() {
	conf := config.GlobalConfig
	var cloudflareDDNSClients []cloudflare.Cloudflare

	for _, target := range conf.Targets {
		cloudflareDDNSClients = append(cloudflareDDNSClients, cloudflare.Cloudflare{
			DNSName: target.DNSName,
			ZoneID:  target.ZoneID,
			Token:   target.Token,
		})
	}
	for _, client := range cloudflareDDNSClients {
		err := RunUpdate(client)
		if err != nil {
			fmt.Printf("Failed to update Cloudflare DNS records: %s\n", err.Error())
		}
	}
}
