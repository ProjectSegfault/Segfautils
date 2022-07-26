package api

import (
	"net/http"
	"html/template"
	"io"
	"log"

    "github.com/kataras/hcaptcha"

	"os"
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
	http.HandleFunc("/api/form", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			io.WriteString(w, "I got your ")
			io.WriteString(w, "\n" + r.RemoteAddr)
		case http.MethodPost:
			io.WriteString(w, "You have no mail to POST!")
		default:
			http.Error(w, "Method isn't allowed!\nYou may only GET or POST here, not " + r.Method, http.StatusMethodNotAllowed)
		}
		log.Println("[HTTP] " + r.RemoteAddr + " accessed /api/form with method " + r.Method)
    })

	http.HandleFunc("/form", renderTestForm)

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
