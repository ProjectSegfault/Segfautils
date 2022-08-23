package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func Port() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting segfautils.port", err.Error())
	}
	result := strconv.Itoa(viper.GetInt("segfautils.port"))
	return result
}
