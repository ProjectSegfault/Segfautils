package config

import (
	"log"

	"github.com/spf13/viper"
)

func OptAnn() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting options.announce", err.Error())
	}
	result := viper.GetString("options.announce")
	return result
}
