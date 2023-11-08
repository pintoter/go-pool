package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "./config/config.yml"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var config = new(Config)
var once sync.Once

func Get() *Config {
	once.Do(func() {
		log.Println("reading app configuration")

		err := cleanenv.ReadConfig(configPath, config)
		if err != nil {
			log.Fatal(err)
		}
	})
	return config
}
