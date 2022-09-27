package api

import (
	"errors"
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
	resAnn    = config.OptAnn()
)

func AnnCheck() {
	if resAnn == "false" {
		log.Println("[Segfautils] ℹ Announcements are disabled")
		http.HandleFunc("/announcements", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Announcements are disabled.", http.StatusServiceUnavailable)
		})
		http.HandleFunc("/api/set/announcements", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "{\"enabled\": \"false\"}", http.StatusOK)
		})
	} else {
		AnnPage()
		http.HandleFunc("/api/set/announcements", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "{\"enabled\": \"true\"}", http.StatusOK)
		})
		Announcements()
	}
}

func AnnPage() {
	http.HandleFunc("/announcements", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/announcements.html")
	})
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
