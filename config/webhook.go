package config

import (
	"log"

	"github.com/spf13/viper"
)

func ShoutrrrURL() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config. Error getting: segfautils.shoutrrr_url", err.Error())
	}
	result := viper.GetString("segfautils.shoutrrr_url")
	return result
}
