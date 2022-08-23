package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func HCaptchaSecretKey() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config for getting hcaptcha.secret_key", err.Error())
	}
	result := strconv.Itoa(viper.GetInt("hcaptcha.secret_key"))
	return result
}
