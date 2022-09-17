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
	Port string
}

func main() {
	log.Println("[Segfautils] Starting")
	utils.CheckConfig()

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := StaticThingy{
			Port: config.Port(),
		}
		tmpl.Execute(w, data)
	})

	log.Println("[HTTP] Starting server")
	api.CheckAnn()
	api.FormCheck()

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "welcome to hell")
	})

	log.Println("[HTTP] HTTP server is now running at " + config.Port() + "!")
	log.Println(http.ListenAndServe(":"+config.Port(), nil))
}
