package config

import (
	"log"

	"github.com/spf13/viper"
)

func WebhookURL() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting segfautils.webhook_url", err.Error())
	}
	result := viper.GetString("segfautils.webhook_url")
	return result
}
