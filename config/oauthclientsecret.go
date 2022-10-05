package config

import (
	"log"

	"github.com/spf13/viper"
)

func OAuthClientSecret() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting oauth.client_secret", err.Error())
	}
	result := viper.GetString("oauth.client_secret")
	return result
}
