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
	tray.InitTray()
}
