package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type ApiConfig struct {
	Env    string
	Port   int
	Twilio Twilio
	DB     DB
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	SSLMode  string
}
type Twilio struct {
	AccountSid  string
	AuthToken   string
	PhoneNumber string
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
		Twilio: Twilio{
			AccountSid:  viper.GetString("TWILIO_ACCOUNT_SID"),
			AuthToken:   viper.GetString("TWILIO_AUTH_TOKEN"),
			PhoneNumber: viper.GetString("TWILIO_PHONE_NUMBER"),
		},
	}

	fmt.Printf("Configuration loaded: %+v\n", conf)
	return conf
}
