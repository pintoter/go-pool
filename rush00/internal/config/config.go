package config

import (
	"log"
	"rush00/pkg/database/postgres"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const configPath = "./config/config.yml"

type Config struct {
	Client struct {
		Address string `yaml:"address"`
	} `yaml:"Client"`
	DB postgres.DBConfig `yaml:"DB"`
}

var config = new(Config)
var once sync.Once

func Get() *Config {
	once.Do(func() {
		log.Println("reading configuration")

		err := cleanenv.ReadConfig(configPath, config)
		if err != nil {
			log.Fatal("can't read config file:", err)
		}
	})
	return config
}
