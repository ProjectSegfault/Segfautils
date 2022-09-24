package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/ProjectSegfault/segfautils/config"
	"github.com/ProjectSegfault/segfautils/utils"
	"github.com/kataras/hcaptcha"
)

var (
	siteKey    = config.HCaptchaSiteKey()
	secretKey  = config.HCaptchaSecretKey()
	webhookURL = config.WebhookURL()
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
			http.Error(w, "{\"enabled\": \"false\"}", http.StatusServiceUnavailable)
		})

	} else {
		FormPage()
		http.HandleFunc("/api/set/form", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "{\"enabled\": \"true\"}", http.StatusServiceUnavailable)
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
			postData := url.Values{
				"content": {"IP " + utils.GetUserIP(r) + "failed captcha!\nhttps://abuseipdb.com/check/" + utils.GetUserIP(r)},
			}
			req, err := http.PostForm(webhookURL, postData)
			if err != nil {
				log.Fatal("Something went terribly wrong!", err)
			}

			fmt.Fprint(io.Discard, req) // I don't want the result of the request in stdout
		} else {
			fmt.Fprintf(w, "Thanks for your message, and thanks for doing the captcha!\nPlease ignore how different this page looks to the page you were on earlier. I'll figure it out eventually!\n%#+v", hcaptchaResp)
			postData := url.Values{
				"content": {"IP " + utils.GetUserIP(r) + "\nFrom " + r.FormValue("email") + " with feedback type " + r.FormValue("commentType") + ":\n" + "**" + r.FormValue("message") + "**\n https://abuseipdb.com/check/" + utils.GetUserIP(r)},
			}
			if r.FormValue("webhook") != "" {
				fmt.Fprintf(w, "\nThanks for trying Segfautils Contact Form :)")
				postData := url.Values{
					"content": {"**Note: you are currently testing our form example. Please check out the actual project at https://github.com/ProjectSegfault/segfautils if you found this neat! It's not hard to self-host :)**\n" + "IP " + utils.GetUserIP(r) + "\nFrom " + r.FormValue("email") + " with feedback type " + r.FormValue("commentType") + ":\n" + "**" + r.FormValue("message") + "**\n https://abuseipdb.com/check/" + utils.GetUserIP(r)},
				}
				req, err := http.PostForm(r.FormValue("webhook"), postData)
				if err != nil {
					log.Println("Someone tried to send a webhook, but it failed!")
				}
				fmt.Fprint(io.Discard, req) // I don't want the result of the demo request in stdout at ALL.
			} else {
				req, err := http.PostForm(webhookURL, postData)
				if err != nil {
					log.Fatal("Something went terribly wrong!", err)
				}
				fmt.Fprint(io.Discard, req) // Out with your request! I don't want it.
			}
		}
	default:
		http.Error(w, "Method isn't allowed!\nYou may only POST here, not "+r.Method, http.StatusMethodNotAllowed)
	}
	log.Println("[HTTP] " + utils.GetUserIP(r) + " accessed /api/form with method " + r.Method)
}
