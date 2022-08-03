package api

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/feeds"
)

var (
	authToken = os.Getenv("SEGFAUTILS_AUTHTOKEN")
)

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
		if r.FormValue("title") == "" || r.FormValue("link") == "" || r.FormValue("severity") == "" {
			http.Error(w, "Your request is not proper. Please add a title, link, and severity.", http.StatusBadRequest)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			now := time.Now()
			feed := &feeds.Feed{
				Title:       r.FormValue("title"),
				Link:        &feeds.Link{Href: r.FormValue("link")},
				Description: r.FormValue("severity"),
				Created:     now,
			}

			json, err := feed.ToJSON()
			if err != nil {
				log.Fatal(err)
			}

			f, err := os.Create("./static/announcements.json")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(json)

			if err2 != nil {
				log.Fatal(err2)
			}

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
		if _, err := os.Stat("./static/announcements.json"); errors.Is(err, os.ErrNotExist) {
			http.Error(w, "If you're gonna delete the annoucement, there has to be an announcement in the first place.", http.StatusNotFound)
			return
		} else {
			err := os.Remove("./static/announcements.json")
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
	if _, err := os.Stat("./static/announcements.json"); errors.Is(err, os.ErrNotExist) {
		http.Error(w, "There are no announcements.", http.StatusNotFound)
		return
	} else {
		f, err := os.Open("./static/announcements.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}
}
