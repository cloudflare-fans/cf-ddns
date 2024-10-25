package main

import (
	"cf-ddns/cloudflare"
	"cf-ddns/tray"
)

var configFilePath = "config.yaml"

func init() {
	err := cloudflare.GlobalConfig.InitConfig(configFilePath)
	if err != nil {
		panic(err)
	}
	err = cloudflare.DDNSScheduleConfigs.InitializeSchedule()
	if err != nil {
		panic(err)
	}
}

func main() {
	tray.InitTray()
}
