package config

import (
	"log"

	"github.com/spf13/viper"
)

func OptForm() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config. Error getting: options.form", err.Error())
	}
	result := viper.GetString("options.form")
	return result
}
