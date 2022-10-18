// Please ignore my terrible code :) It works
package utils

import (
	"log"

	"github.com/ProjectSegfault/segfautils/config"
)

func CheckConfig() {
	if config.Port() == "0" {
		log.Fatal("[Segfautils] ❌ You need to set the port you'd like to use in the config file. Check documentation for more information.")
	} else {
		log.Println("[Segfautils] ✅ segfautils.port is set to", config.Port())
	}
	if config.AuthToken() == "YOURAUTHTOKEN" || config.AuthToken() == "" {
		log.Fatal("[Segfautils] ❌ You need to set the authentication token you'd like to use in the config file. Check documentation for more information.")
	} else {
		log.Println("[Segfautils] ✅ segfautils.auth_token is set!")
	}
	if config.ShoutrrrURL() == "" {
		log.Fatal("[Segfautils] ❌ You need to set the Webhook URL you'd like to use in the config file. Check documentation for more information.")
	} else {
		log.Println("[Segfautils] ✅ segfautils.shoutrrr_url is set!")
	}
	// Hcaptcha stuff
	if config.HCaptchaSecretKey() == "YOURSECRETKEY" || config.HCaptchaSecretKey() == "" {
		log.Fatal("[Segfautils] ❌ You need to set the HCaptcha secret you'd like to use in the config file. Check documentation for more information.")
	} else {
		log.Println("[Segfautils] ✅ segfautils.hcaptcha_secret is set!")
	}
	if config.HCaptchaSiteKey() == "YOURSITEKEY" || config.HCaptchaSiteKey() == "" {
		log.Println("[Segfautils] ⚠️ The HCaptcha site key isn't set. You don't have to, but the demo form will not work without it. Check documentation for more information.")
	} else {
		log.Println("[Segfautils] ✅ hcaptcha.site_key is set!")
	}
	log.Println("[Segfautils] ✅ All config checks passed!")
}
