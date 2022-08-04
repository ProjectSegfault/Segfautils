package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ProjectSegfault/segfautils/api"
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
	utils.CheckEnv()
	log.Println("[HTTP] Starting server")
	port := os.Getenv("SEGFAUTILS_PORT")
	hcaptcha_site_key := os.Getenv("HCAPTCHA_SITE_KEY")
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := StaticThingy{
			Port: port,
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
	api.Form()
	api.Announcements()
	log.Println("[HTTP] HTTP server is now running at " + port + "!")
	log.Println(http.ListenAndServe(":"+port, nil))
}
