// Please ignore my terrible code :) It works
package otherthings

import (
	"os"
	"log"
)

var (
	unused string
	ok1 bool
)

func CheckEnv() {
	unused, ok1 = os.LookupEnv("SEGFAUTILITIES_PORT")
	if ok1 {
		log.Println("[Segfautilities] Environment variable SEGFAUTILITIES_PORT is set as " + unused)
	} else {
		log.Fatal("[Segfautilities] Environment variable SEGFAUTILITIES_PORT is not set! Please set it to a number, for example 6893")
	}
	unused, ok1 = os.LookupEnv("HCAPTCHA_SITE_KEY")
	if !ok1 {
		log.Fatal("[Segfautilities] Environment variable HCAPTCHA_SITE_KEY is not set! Please set it to the site key you got from hCaptcha.")
	} else {
		log.Println("[Segfautilities] Environment variable HCAPTCHA_SITE_KEY is set as " + unused)
	}
	unused, ok1 = os.LookupEnv("HCAPTCHA_SECRET_KEY")
	if !ok1 {
		log.Fatal("[Segfautilities] Environment variable HCAPTCHA_SECRET_KEY is not set! Please set it to the secret key you got from hCaptcha.")
	} else {
		log.Println("[Segfautilities] Environment variable HCAPTCHA_SECRET_KEY is set!")
	}
	log.Println("[Segfautilities] ✅ Passed the Environment Variables check")
}