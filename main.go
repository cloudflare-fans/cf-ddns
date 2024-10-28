package main

import (
	"github.com/cloudflare-fans/cf-ddns/cloudflare"
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
}
