package main

import (
        "net/http"
		"html/template"
		"io"
        "log"
        "github.com/ProjectSegfault/segfautilities/otherthings"
        "os"
        "github.com/ProjectSegfault/segfautilities/api"
)

type StaticThingy struct {
    Port string
}

var port string
var shit bool

func main() {
    log.Println("[Segfautilities] Starting")
    otherthings.CheckEnv()
    log.Println("[HTTP] Starting server")
    port := os.Getenv("SEGFAUTILITIES_PORT")
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
            Port: port,
        }
        tmpl_form.Execute(w, data)
    })

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "welcome to hell")
    })
    api.Form()
    log.Println("[HTTP] HTTP server is now running at " + port + "!")
    log.Println(http.ListenAndServe(":" + port, nil))
}