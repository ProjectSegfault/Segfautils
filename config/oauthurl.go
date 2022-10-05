package config

import (
	"log"

	"github.com/spf13/viper"
)

func OAuthURL() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting oauth.auth_url", err.Error())
	}
	result := viper.GetString("oauth.auth_url")
	return result
}
