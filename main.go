package main

import (
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/ProjectSegfault/segfautils/api"
	"github.com/ProjectSegfault/segfautils/config"
	"github.com/ProjectSegfault/segfautils/utils"
)

type StaticThingy struct {
	Port            string
	HCaptchaSiteKey string
}

var port string
var shit bool

func main() {
	log.Println("[Segfautils] Starting")
	utils.CheckConfig()
	log.Println("[HTTP] Starting server")
	hcaptcha_site_key := config.HCaptchaSiteKey()
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := StaticThingy{
			Port: config.Port(),
		}
		tmpl.Execute(w, data)
	})

	tmpl_form := template.Must(template.ParseFiles("static/form.html"))
	http.HandleFunc("/form/", func(w http.ResponseWriter, r *http.Request) {
		data := StaticThingy{
			HCaptchaSiteKey: hcaptcha_site_key,
		}
		tmpl_form.Execute(w, data)
	})

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "welcome to hell")
	})
	http.HandleFunc("/announcements", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/announcements.html")
	})
	api.FormCheck()
	api.CheckAnn()
	log.Println("[HTTP] HTTP server is now running at " + config.Port() + "!")
	log.Println(http.ListenAndServe(":"+config.Port(), nil))
}
