package main

import (
	"github.com/cloudflare-fans/cf-ddns/cloudflare"
	"github.com/cloudflare-fans/cf-ddns/sys_conf"
	"github.com/cloudflare-fans/cf-ddns/web_server"
)

var configFilePath = "config.yaml"
var configSystemFilePath = "config-system.yaml"

func init() {
	err := cloudflare.GlobalConfig.InitConfig(configFilePath)
	if err != nil {
		panic(err)
	}
	err = cloudflare.DDNSScheduleConfigs.InitializeSchedule()
	if err != nil {
		panic(err)
	}
	err = sys_conf.InitSysConf(configSystemFilePath)
	if err != nil {
		panic(err)
	}
}

func main() {
	server := web_server.Init()
	server.RegisterRouter()
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
