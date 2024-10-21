package main

import (
	"cf-ddns/config"
	"cf-ddns/tray"
)

var configFilePath = "config.yaml"

func init() {
	config.InitConfig(configFilePath)
}

func main() {
	/*
		conf := config.GlobalConfig

		for _, target := range conf.Targets {
			client := cloudflare.NewCFClient(target.DNSName, target.ZoneID, target.Token)
			err := client.RunDDNS()
			if err != nil {
				fmt.Printf("Failed to update Cloudflare DNS records: %s\n", err.Error())
			}
		}
	*/

	tray.InitTray()
}
