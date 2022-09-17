package api

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ProjectSegfault/segfautils/config"
	"github.com/goccy/go-json"
)

var (
	announcements = config.OptAnn()
	form          = config.OptForm()
)

func Settings() {
	CheckSet()
	http.HandleFunc("/api/options", getOpt)
}

func CheckSet() {
	os.Remove("./data/options.json")
	if form == "true" && announcements == "false" {
		data := map[string]interface{}{
			"Announcements": "false",
			"Form":          "true",
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Could not marshal json : %s\n", err)
			return
		}

		ioutil.WriteFile("./data/options.json", jsonData, os.ModePerm)

	} else if form == "true" && announcements == "true" {
		data := map[string]interface{}{
			"Announcements": "true",
			"Form":          "true",
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Could not marshal json : %s\n", err)
			return
		}

		ioutil.WriteFile("./data/options.json", jsonData, os.ModePerm)

	} else if form == "false" && announcements == "true" {
		data := map[string]interface{}{
			"Announcements": "true",
			"Form":          "false",
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Could not marshal json : %s\n", err)
			return
		}

		ioutil.WriteFile("./data/options.json", jsonData, os.ModePerm)

	} else {
		resp := []byte("The fuck do you want me to do then?")
		ioutil.WriteFile("./data/options.json", resp, os.ModePerm)
	}
}
func getOpt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if _, err := os.Stat("./data/options.json"); errors.Is(err, os.ErrNotExist) {
		http.Error(w, "There is nothing to see here.", http.StatusNotFound)
		return
	} else {
		f, err := os.Open("./data/options.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}
}
