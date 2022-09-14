package config

import (
	"log"

	"github.com/spf13/viper"
)

func OAuthClientID() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting oauth.client_id", err.Error())
	}
	result := viper.GetString("oauth.client_id")
	return result
}
