package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "config.yml"
)

type Config struct {
	LocalAddress string `yaml:"LocalAddress"`
	LocalPort    int    `yaml:"LocalPort"`
	Network      int    `yaml:"Network"`
}

func (c *Config) LoadConfig() *Config {
	yml, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Could not load config file. Error: %v\r\n", err)
	}

	err = yaml.Unmarshal(yml, c)
	if err != nil {
		log.Fatalf("Could not load config file. Error: %v\r\n", err)
	}

	return c
}
