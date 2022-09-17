package api

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ProjectSegfault/segfautils/config"
	"github.com/goccy/go-json"
)

var (
	authToken = config.AuthToken()
)

func CheckAnn() {
	jsonFile, err := os.Open("./data/options.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	res := result["Announcements"]
	if res == "true" {
		Announcements()
	} else {
		log.Println("Announcements disabled")
	}
}

func Announcements() {
	http.HandleFunc("/api/announcements", getAnnouncements)
	http.HandleFunc("/api/announcements/post", handleAnnouncements)
	http.HandleFunc("/api/announcements/delete", handleAnnouncementDeleteRequest)
}

func handleAnnouncements(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.FormValue("token") != authToken {
		http.Error(w, "You need to provide the authorization token given to you by your system administrator in order to post an announcement.", http.StatusUnauthorized)
		return
	} else {
		if r.FormValue("title") == "" || r.FormValue("severity") == "" {
			http.Error(w, "Your request is not proper. Please add a title and severity.", http.StatusBadRequest)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			now := time.Now().Unix()
			data := map[string]interface{}{
				"title":    r.FormValue("title"),
				"link":     r.FormValue("link"),
				"severity": r.FormValue("severity"),
				"created":  now,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Printf("could not marshal json: %s\n", err)
				return
			}

			ioutil.WriteFile("./data/announcements.json", jsonData, os.ModePerm)

			w.Write([]byte("Announcement posted!"))
		}
		return
	}
}

func handleAnnouncementDeleteRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.FormValue("token") != authToken {
		http.Error(w, "You need to provide the authorization token given to you by your system administrator in order to delete an announcement.", http.StatusUnauthorized)
		return
	} else {
		if _, err := os.Stat("./data/announcements.json"); errors.Is(err, os.ErrNotExist) {
			http.Error(w, "If you're gonna delete the annoucement, there has to be an announcement in the first place.", http.StatusNotFound)
			return
		} else {
			err := os.Remove("./data/announcements.json")
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Announcement deleted!"))
			return
		}
	}
}

func getAnnouncements(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if _, err := os.Stat("./data/announcements.json"); errors.Is(err, os.ErrNotExist) {
		http.Error(w, "There are no announcements.", http.StatusNotFound)
		return
	} else {
		f, err := os.Open("./data/announcements.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}
}
