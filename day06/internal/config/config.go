package config

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type HTTPConfig struct {
	Host string
	Port string
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Sslmode  string
}

func (db *DB) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db.User, db.Password, db.Host, db.Port, db.Name, db.Sslmode)
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
func ReadConfigTxt(configTxt string) (*Config, error) {
	file, err := os.Open(configTxt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		credentials := strings.Split(line, ":")

		switch credentials[0] {
		case "admin":
			cfg.AdminCredentials.Login = credentials[1]
		case "password":
			cfg.AdminCredentials.Password = credentials[1]
		case "dbuser":
			cfg.DB.User = credentials[1]
		case "dbpassword":
			cfg.DB.Password = credentials[1]
		case "dbhost":
			cfg.DB.Host = credentials[1]
		case "dbport":
			cfg.DB.Port = credentials[1]
		case "dbname":
			cfg.DB.Name = credentials[1]
		case "sslmode":
			cfg.DB.Sslmode = credentials[1]
		case "host":
			cfg.Addr.Host = credentials[1]
		case "port":
			cfg.Addr.Port = credentials[1]
		default:
			return nil, errors.New("undefined credential")
		}
	}

	return cfg, nil
}
