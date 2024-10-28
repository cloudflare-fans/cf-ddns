package sys_conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type sysConf struct {
	Server struct {
		Listen string `yaml:"listen"`
	} `yaml:"server"`
}

var GlobalSystemConf sysConf

func InitSysConf(configFilePath string) (err error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &GlobalSystemConf)
	if err != nil {
		return err
	}
	return nil
}
