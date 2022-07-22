// Please ignore this :))))))) :(
package otherthings

import (
	"os"
	"log"
)

var unused string
var ok1 bool

func CheckEnv() {
	unused, ok1 = os.LookupEnv("SEGFAUTILITIES_PORT")
	if ok1 {
		log.Println("[Segfautilities] Environment variable SEGFAUTILITIES_PORT is set as " + unused)
	} else {
		log.Fatal("[Segfautilities] Environment variable SEGFAUTILITIES_PORT is not set! Please set it to a number, for example 6893")
	}

	log.Println("[Segfautilities] ✅ Passed the Environment Variables check")
}