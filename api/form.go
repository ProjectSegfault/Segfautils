package api

import (
	"net/http"
	"html/template"
	"log"

    "github.com/kataras/hcaptcha"

	"os"
	"fmt"

	"net/url"
	"io"
)

var (
	siteKey   = os.Getenv("HCAPTCHA_SITE_KEY")
	secretKey = os.Getenv("HCAPTCHA_SECRET_KEY")
)

var (
	client       = hcaptcha.New(secretKey) /* See `Client.FailureHandler` too. */
	testForm = template.Must(template.ParseFiles("./static/testform.html"))
)

func Form() {
	http.HandleFunc("/api/form", client.HandlerFunc(theActualFormCode))

	http.HandleFunc("/form", renderTestForm)

}

func theActualFormCode(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			hcaptchaResp, ok := hcaptcha.Get(r)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Seems like captcha failed, you didn't complete the captcha or you are a bot. Please try again.\nPlease note that your IP has been logged in our systems for manual review to check if you're an abusive user. If you're seen as abusive, you will be blacklisted.")
				postData := url.Values{
					"content": {"IP " + r.RemoteAddr + "failed captcha!"},
				}
				req, err := http.PostForm("https://canary.discord.com/api/webhooks/1001546472971587674/XpxdpvlrEvEo7jnRmWyrEudTW0eNC_LpNVd-uC20R44e_efgwvhzSNQ26IaGeO9yan6f", postData)
				if err != nil {
     			   log.Fatal(err)
    			}

				fmt.Fprint(io.Discard, req) // I don't want the result of the request in stdout
				return
			}
			log.Printf(r.FormValue("aisdjaoidjoiajsidaoj"))
			fmt.Fprintf(w, "Thanks for your message, and thanks for doing the captcha!\n%#+v", hcaptchaResp)
			postData := url.Values{
				"content": {"IP " + r.RemoteAddr + "\nFrom " + r.FormValue("email") + " with feedback type " + r.FormValue("commentType") + ":\n" + "**" + r.FormValue("message") + "**"},
			}
			req, err := http.PostForm("https://discord.com/api/webhooks/1001546472971587674/XpxdpvlrEvEo7jnRmWyrEudTW0eNC_LpNVd-uC20R44e_efgwvhzSNQ26IaGeO9yan6f", postData)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprint(io.Discard, req) // Out with your request! I don't want it.
		default:
			http.Error(w, "Method isn't allowed!\nYou may only POST here, not " + r.Method, http.StatusMethodNotAllowed)
		}
		log.Println("[HTTP] " + r.RemoteAddr + " accessed /api/form with method " + r.Method)
}

func renderTestForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	testForm.Execute(w, map[string]string{
		"SiteKey": siteKey,
	})
}

// testForm is only used in development. I will remove it when I've added it to the website
// Oh also, you need to add the following to your hosts file:
// 127.0.0.1 epicwebsite.com
// and visit epicwebsite.com:(yourport)/form. hCaptcha doesn't work in localhost unfortunately :(
