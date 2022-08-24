package config

import (
	"log"

	"github.com/spf13/viper"
)

func HCaptchaSecretKey() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting hcaptcha.secret_key", err.Error())
	}
	result := viper.GetString("hcaptcha.secret_key")
	return result
}
