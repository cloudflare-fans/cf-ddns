package main

import (
	"cf-ddns/cloudflare"
	"cf-ddns/tray"
)

var configFilePath = "config.yaml"

func init() {
	cloudflare.GlobalConfig.InitConfig(configFilePath)
	err := cloudflare.DDNSScheduleConfigs.InitializeSchedule()
	if err != nil {
		panic(err)
	}
}

func main() {
	tray.InitTray()
}
