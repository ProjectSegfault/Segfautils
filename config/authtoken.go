package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	result string
)

func AuthToken() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting segfautils.auth_token", err.Error())
	}
	result := viper.GetString("segfautils.auth_token")
	return result
}
