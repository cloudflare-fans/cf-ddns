package main

import (
	"github.com/cloudflare-fans/cf-ddns/cloudflare"
	"github.com/cloudflare-fans/cf-ddns/sys_conf"
	"github.com/cloudflare-fans/cf-ddns/web_server"
)

var configFilePath = "config.yaml"
var configSystemFilePath = "config-system.yaml"

func init() {
	// init cloudflare DDNS tasks configs
	err := cloudflare.GlobalConfig.InitConfig(configFilePath)
	if err != nil {
		panic(err)
	}
	// initialize DDNS scheduling tasks
	err = cloudflare.DDNSScheduleConfigs.InitializeSchedule()
	if err != nil {
		panic(err)
	}
	// init system configs
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
