package config

import (
	"log"
	"sync"
)

const (
	configPath = "./config/config.yml"
)

type Config struct {
	BindIP string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

var instance *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		log.Println("reading app configuration")
		instance = &Config{}

		err := cleanenv. // dowload cleanenv
	})
	return &Config{}
}
