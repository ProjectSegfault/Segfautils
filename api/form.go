package api

import (
	"net/http"
	"log"

    "github.com/kataras/hcaptcha"

	"os"
	"fmt"

	"net/url"
	"io"

    "github.com/ProjectSegfault/segfautilities/otherthings"
)

var (
	siteKey   = os.Getenv("HCAPTCHA_SITE_KEY")
	secretKey = os.Getenv("HCAPTCHA_SECRET_KEY")
	webhookURL = os.Getenv("SEGFAUTILITIES_WEBHOOK_URL")
	client   = hcaptcha.New(secretKey) /* See `Client.FailureHandler` too. */
)

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
					"content": {"IP " + otherthings.GetUserIP(r) + "failed captcha!\nhttps://abuseipdb.com/check/" + otherthings.GetUserIP(r)},
				}
				req, err := http.PostForm(webhookURL, postData)
				if err != nil {
     			   log.Fatal("Something went terribly wrong!", err)
    			}

				fmt.Fprint(io.Discard, req) // I don't want the result of the request in stdout
			} else {
				fmt.Fprintf(w, "Thanks for your message, and thanks for doing the captcha!\nPlease ignore how different this page looks to the page you were on earlier. I'll figure it out eventually!\n%#+v", hcaptchaResp)
				postData := url.Values{
					"content": {"IP " + otherthings.GetUserIP(r) + "\nFrom " + r.FormValue("email") + " with feedback type " + r.FormValue("commentType") + ":\n" + "**" + r.FormValue("message") + "**\n https://abuseipdb.com/check/" + otherthings.GetUserIP(r)},
				}
				if r.FormValue("webhook") != "" {
					fmt.Fprintf(w, "\nThanks for trying Segfautilities Contact Form :)")
					postData := url.Values{
						"content": {"**Note: you are currently testing our form example. Please check out the actual project at https://github.com/ProjectSegfault/segfautilities! It's not hard to self-host :)**\n" + "IP " + otherthings.GetUserIP(r) + "\nFrom " + r.FormValue("email") + " with feedback type " + r.FormValue("commentType") + ":\n" + "**" + r.FormValue("message") + "**\n https://abuseipdb.com/check/" + otherthings.GetUserIP(r)},
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
			http.Error(w, "Method isn't allowed!\nYou may only POST here, not " + r.Method, http.StatusMethodNotAllowed)
		}
		log.Println("[HTTP] " + otherthings.GetUserIP(r) + " accessed /api/form with method " + r.Method)
}