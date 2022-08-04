// Please ignore my terrible code :) It works
package utils

import (
	"log"
	"os"
)

var (
	unused string
	ok1    bool
)

func CheckEnv() {
	unused, ok1 = os.LookupEnv("SEGFAUTILS_PORT")
	if ok1 {
		log.Println("[Segfautils] Environment variable SEGFAUTILS_PORT is set as " + unused)
	} else {
		log.Fatal("[Segfautils] Environment variable SEGFAUTILS_PORT is not set! Please set it to a number, for example 6893")
	}
	unused, ok1 = os.LookupEnv("HCAPTCHA_SITE_KEY")
	if !ok1 || unused == "YOURSITEKEY" {
		log.Println("[Segfautils] Environment variable HCAPTCHA_SITE_KEY is not not set. It isn't required to be set, but without it, the example form will not work.")
	} else {
		log.Println("[Segfautils] Environment variable HCAPTCHA_SITE_KEY is set as " + unused)
	}
	unused, ok1 = os.LookupEnv("HCAPTCHA_SECRET_KEY")
	if !ok1 || unused == "YOURSECRETKEY" {
		log.Fatal("[Segfautils] Environment variable HCAPTCHA_SECRET_KEY is not set! Please set it to the secret key you got from hCaptcha.")
	} else {
		log.Println("[Segfautils] Environment variable HCAPTCHA_SECRET_KEY is set!")
	}
	unused, ok1 = os.LookupEnv("SEGFAUTILS_WEBHOOK_URL")
	if !ok1 || unused == "YOURWEBHOOKURL" {
		log.Fatal("[Segfautils] Environment variable SEGFAUTILS_WEBHOOK_URL is not set! Please set it to your webhook URL. If that URL doesn't work, make an issue on GitHub!")
	} else {
		log.Println("[Segfautils] Environment variable SEGFAUTILS_WEBHOOK_URL is set!")
	}
	unused, ok1 = os.LookupEnv("SEGFAUTILS_AUTHTOKEN")
	if !ok1 || unused == "YOURAUTHTOKEN" {
		log.Fatal("[Segfautils] Environment variable SEGFAUTILS_AUTHTOKEN is not set! Please set it to a token you'd like to use for authorizing actions like announcements.")
	} else {
		log.Println("[Segfautils] Environment variable SEGFAUTILS_AUTHTOKEN is set!")
	}
	log.Println("[Segfautils] ✅ Passed the Environment Variables check")
}
