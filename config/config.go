package config

import (
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type cloudflareTarget struct {
	DNSName     string `json:"dns_name,omitempty" yaml:"dns_name,omitempty"`
	ZoneID      string `json:"zone_id,omitempty" yaml:"zone_id,omitempty"`
	Token       string `json:"token,omitempty" yaml:"token,omitempty"`
	UpdateEvery string `json:"update_every,omitempty" yaml:"update_every,omitempty"`
}

type config struct {
	Targets []cloudflareTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}

var GlobalConfig config

func readConfig(configFilePath string) error {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(configFilePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Config file modified or created, reloading...")
					err := readConfig(configFilePath)
					if err != nil {
						log.Fatal(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// 首次读取配置
	err = readConfig(configFilePath)
	if err != nil {
		log.Fatal("Error reading initial config:", err)
	}
	log.Println("Initial config loaded successfully")
}
