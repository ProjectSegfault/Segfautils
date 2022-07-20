package main

import (
        "net/http"
		"html/template"
		"io"
)

type StaticThingy struct {
    Port string
}

func main() {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := StaticThingy{
            Port: "3000",
        }
        tmpl.Execute(w, data)
    })
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "welcome to hell")
    })
        http.ListenAndServe(":3000", nil)
}