package cloudflare

import (
	"cf-ddns/config"
	"fmt"
)

func DoUpdateCFDDNS(successCallback func(), errorCallback func(error)) {
	conf := config.GlobalConfig

	for _, target := range conf.Targets {
		client := NewCFClient(target.DNSName, target.ZoneID, target.Token)
		err := client.RunDDNS()
		if err != nil {
			fmt.Printf("Failed to update Cloudflare DNS records: %s\n", err.Error())
			errorCallback(err)
		} else {
			successCallback()
		}
	}
}
