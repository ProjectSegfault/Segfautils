package config

import (
	"log"

	"github.com/spf13/viper"
)

func OAuthRedirectURL() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting oauth.redirect_url", err.Error())
	}
	result := viper.GetString("oauth.redirect_url")
	return result
}
