package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func HCaptchaSiteKey() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting hcaptcha.site_key", err.Error())
	}
	result := strconv.Itoa(viper.GetInt("hcaptcha.site_key"))
	return result
}
