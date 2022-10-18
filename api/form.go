package api

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ProjectSegfault/segfautils/config"
	"github.com/ProjectSegfault/segfautils/utils"
	"github.com/kataras/hcaptcha"

	"github.com/containrrr/shoutrrr"
)

var (
	siteKey    = config.HCaptchaSiteKey()
	secretKey  = config.HCaptchaSecretKey()
	webhookURL = config.ShoutrrrURL()
	client     = hcaptcha.New(secretKey) /* See `Client.FailureHandler` too. */
	resForm    = config.OptForm()
)

func FormCheck() {
	if resForm == "false" {
		log.Println("[Segfautils] ℹ Contact form is disabled")
		http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Form is disabled.", http.StatusServiceUnavailable)
		})
		http.HandleFunc("/api/set/form", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "{\"enabled\": \"false\"}", http.StatusOK)
		})

	} else {
		FormPage()
		http.HandleFunc("/api/set/form", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "{\"enabled\": \"true\"}", http.StatusOK)
		})
		Form()
	}
}

func FormPage() {
	type StaticThing struct {
		HCaptchaSiteKey string
	}

	tmpl_form := template.Must(template.ParseFiles("static/form.html"))
	http.HandleFunc("/form/", func(w http.ResponseWriter, r *http.Request) {

		hcaptcha_site_key := config.HCaptchaSiteKey()
		data := StaticThing{
			HCaptchaSiteKey: hcaptcha_site_key,
		}
		tmpl_form.Execute(w, data)
	})
}

func Form() {
	http.HandleFunc("/api/form", client.HandlerFunc(theActualFormCode))
}

func theActualFormCode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		hcaptchaResp, ok := hcaptcha.Get(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Seems like captcha failed, you didn't complete the captcha or you are a bot. Please try again.\nPlease note that your IP has been logged in our systems for manual review to check if you're an abusive user. If you're seen as abusive, you will be blacklisted.\nYour message has not been sent.")
			webhook := shoutrrr.Send(config.ShoutrrrURL(), "IP "+utils.GetUserIP(r)+" failed to complete the captcha.\nhttps://abuseipdb.com/check/"+utils.GetUserIP(r))
			if webhook != nil {
				log.Println("[Segfautils] ✖ Failed to send webhook")
			}
		} else {
			fmt.Fprintf(w, "Thanks for your message, and thanks for doing the captcha!\nPlease ignore how different this page looks to the page you were on earlier. I'll figure it out eventually!\n%#+v", hcaptchaResp)
			shoutrrr := shoutrrr.Send(config.ShoutrrrURL(), "IP "+utils.GetUserIP(r)+"\nFrom "+r.FormValue("email")+" with feedback type "+r.FormValue("commentType")+":\n"+"**"+r.FormValue("message")+"**\n https://abuseipdb.com/check/"+utils.GetUserIP(r))
			if shoutrrr != nil {
				log.Fatal("Something went terribly wrong!", shoutrrr)
			}
		}
	default:
		http.Error(w, "Method isn't allowed!\nYou may only POST here, not "+r.Method, http.StatusMethodNotAllowed)
	}
	log.Println("[HTTP] " + utils.GetUserIP(r) + " accessed /api/form with method " + r.Method)
}
