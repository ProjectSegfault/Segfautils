package api

import (
	"net/http"
	"io"
	"log"
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
}