// Please ignore this :))))))) :(
package main

import (
	"os"
	"fmt"
)

func main() {
	LookedupEnv := os.LookupEnv("SEGFAUTILITiES_PORT")
	if LookedupEnv === false {
		fmt.Println("haha")
	}
	fmt.Println(os.Getenv("SEGFAUTILITIES_PORT"))
}