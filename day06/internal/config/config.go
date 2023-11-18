package config

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

type HTTPConfig struct {
	Host string
	Port string
}

type DB struct {
	DSN string
}

func (db *DB) GetDSN() string {
	return db.DSN
}

type AdminCredentials struct {
	Login    string
	Password string
}

type Config struct {
	Addr             HTTPConfig
	AdminCredentials AdminCredentials
	DB               DB
}

// Read configurations from file and init instance Config
func ReadConfigTxt(configTxt string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configTxt)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Println(err)
				return err
			}
		}

		credentials := strings.Split(line, ":")

		switch credentials[0] {
		case "admin":
			cfg.AdminCredentials.Login = credentials[1]
		case "password":
			cfg.AdminCredentials.Password = credentials[1]
		case "dsn":
			cfg.DB.DSN = credentials[1]
		case "host":
			cfg.Addr.Host = credentials[1]
		case "port":
			cfg.Addr.Port = credentials[1]
		default:
			return errors.New("undefined credential")
		}
	}

	return nil
}
