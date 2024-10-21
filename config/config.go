package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type cloudflareTarget struct {
	DNSName string `json:"dns_name,omitempty" yaml:"dns_name,omitempty"`
	ZoneID  string `json:"zone_id,omitempty" yaml:"zone_id,omitempty"`
	Token   string `json:"token,omitempty" yaml:"token,omitempty"`
}

type config struct {
	Targets []cloudflareTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}

var GlobalConfig config

func InitConfig(configFilePath string) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
		return
	}
	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		log.Fatalf("无法应用配置文件: %v", err)
		return
	}
}
