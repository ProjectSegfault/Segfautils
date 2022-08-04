package utils

import (
	"net/http"
)

// Thanks random StackOverflow answerer

func GetUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-REAL-IP")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-FORWARDED-FOR")
	}
	return IPAddress
}
