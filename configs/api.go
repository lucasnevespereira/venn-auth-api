package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type ApiConfig struct {
	Env  string
	Port int
	DB   DB
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	SSLMode  string
}

func LoadApiConfig() ApiConfig {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	conf := ApiConfig{
		Env:  viper.GetString("ENV"),
		Port: viper.GetInt("PORT"),
		DB: DB{
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			Name:     viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSL_MODE"),
		},
	}

	fmt.Printf("Configuration loaded: %+v\n", conf)
	return conf
}
