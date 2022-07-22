package main

import (
        "net/http"
		"html/template"
		"io"
        "log"
        "github.com/ProjectSegfault/segfautilities/otherthings"
        "os"
)

type StaticThingy struct {
    Port string
}

var pieceof string
var shit bool

func main() {
    log.Println("[Segfautilities] Starting")
    otherthings.CheckEnv()
    log.Println("[HTTP] Starting server")
    pieceof := os.Getenv("SEGFAUTILITIES_PORT") // I hate this
	tmpl := template.Must(template.ParseFiles("static/index.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := StaticThingy{
            Port: pieceof,
        }
        tmpl.Execute(w, data)
    })
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "welcome to hell")
    })
    log.Println("[HTTP] HTTP server is now running at " + pieceof + "!")
    log.Println(http.ListenAndServe(":" + pieceof, nil))
}