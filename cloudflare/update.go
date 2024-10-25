package cloudflare

import (
	"fmt"
)

func DoUpdateCFDDNS(successCallback func(), errorCallback func(error)) {
	conf := GlobalConfig

	for _, target := range conf.Targets {
		err := target.RunDDNS()
		if err != nil {
			fmt.Printf("Failed to update Cloudflare DNS records: %s\n", err.Error())
			errorCallback(err)
		} else {
			successCallback()
		}
	}
}
