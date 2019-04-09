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
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yml, c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
