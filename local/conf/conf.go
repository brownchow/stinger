package conf

import (
	"github.com/ritterhou/stinger/core/common"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	HttpPort      int      `yaml:"http_port"`
	Global        bool     `yaml:"global"`
	Domains       []string `yaml:"domains"`
	LocalPort     int      `yaml:"local_port"`
	ServerAddress string   `yaml:"server_address"`
	Password      string   `yaml:"password"`
	LogFile       string   `yaml:"log_file"`
	LogLevel      string   `yaml:"log_level"`
}

var conf Conf

func LoadConf(filename string) {
	path := common.GetAbsPath(filename)
	content := common.ReadFile(path)

	err := yaml.Unmarshal(content, &conf)
	if err != nil {
		logrus.Fatal(err)
	}

	if conf.HttpPort == 0 {
		conf.HttpPort = 2600
	}
	if conf.LocalPort == 0 {
		conf.LocalPort = 2680
	}
	if conf.ServerAddress == "" {
		conf.ServerAddress = "127.0.0.1:26800"
	}
	if conf.Password == "" {
		conf.Password = "123456"
	}
	if conf.LogFile == "" {
		conf.LogFile = "stdout"
	}
	if conf.LogLevel == "" {
		conf.LogLevel = "WARN"
	}
}

func GetConf() Conf {
	return conf
}
