package config

import (
	"os"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASS"),
		Host:     "127.0.0.1",
		Port:     "3306",
		Name:     "emission_tracker",
	}
}
