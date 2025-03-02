package internal

import (
	"github.com/mskreczko/uptime-checker/pkg"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Applications         []Application        `yaml:"applications"`
	NotificationSettings NotificationSettings `yaml:"notifications"`
	SMTPSettings         pkg.AWSSesClient     `yaml:"smtp"`
	ListeningPort        int                  `yaml:"listening_port"`
}

func ReadConfig(configPath string) Config {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
