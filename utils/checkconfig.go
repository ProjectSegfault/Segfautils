// Please ignore my terrible code :) It works
package utils

import (
	"fmt"
	"os"
)

var (
	is      = 0
	arrName = [9]string{"Port", "Authentication token", "Webhook URL", "HCaptcha secret", "Hcaptcha site key", "OAuth client ID", "OAuth client secret", "OAuth redirect URL", "OAuth athentication URL"}
	arrCfg  = [9]string{"0", "YOURAUTHTOKEN", "YOURWEBHOOKURL", "YOURSECRETKEY", "YOURSITEKEY", "YOURCLIENTID", "YOURCLIENTSECRET", "YOURREDIRECTURL", "YOURAUTHURL"}
	testCfg = [9]string{"config.Authtoken()", "config.WebhookURL()", "config.HCaptchaSecretKey()", "config.HCaptchaSiteKey()", "config.OAuthClientID()", "config.OAuthClientSecret()", "config.OAuthRedirectURL()", "config.OAuthURL()"}
)

func CheckConfig() {

	for i := 0; i < 8; i++ {
		if testCfg[i] == arrCfg[i] || testCfg[i] == "" {
			fmt.Println(arrCfg[i])
			is = i
			Fail()
		} else {
			fmt.Println(arrCfg[i])
			fmt.Println(testCfg[i])
			is = i
			Check()
		}
	}
}

/*
	func CheckConfig() {
		if config.Port() == "0" {
			Fail()
		} else {
			Check()
		}
		if config.AuthToken() == "YOURAUTHTOKEN" || config.AuthToken() == "" {
			log.Fatal("[Segfautils] ❌ You need to set the authentication token you'd like to use in the config file. Check documentation for more information.")
		} else {
			log.Println("[Segfautils] ✅ segfautils.auth_token is set!")
		}
		if config.WebhookURL() == "YOURWEBHOOKURL" || config.WebhookURL() == "" {
			log.Fatal("[Segfautils] ❌ You need to set the Webhook URL you'd like to use in the config file. Check documentation for more information.")
		} else {
			log.Println("[Segfautils] ✅ segfautils.webhook_url is set!")
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
		log.Println("[segfautils] ✅ All config checks passed!")
		log.Println("[segfautils] ✅ Listening at port ", config.Port())
	}
*/
func Check() {
	fmt.Printf("[Segfautils] ✅ %s is set!\n", arrName[is])
}

func Fail() {
	fmt.Printf("[Segfautils] ❌ You need to set the %s you'd like to use in the config file. Check documentation for more information.\n", arrName[is])
	os.Exit(Error())
}

func OptionalFail() {
	fmt.Printf("[Segfautils] ⚠️ The %s isn't set. You don't have to, but the program will not be able to behave correctly. Check documentation for more information.\n", arrName[is])
}

func Error() int {
	return 12
}
