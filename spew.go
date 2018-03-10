package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// init
	fmt.Println("Spewing...")

	// settings
	spewOut := "spew"
	if os.Getenv("SPEW_OUT") != "" {
		spewOut = os.Getenv("SPEW_OUT")
	}
	spewInterval := "5s"
	if os.Getenv("SPEW_INTERVAL") != "" {
		spewInterval = os.Getenv("SPEW_INTERVAL")
	}
	spewTimespan, err := time.ParseDuration(spewInterval)
	if err != nil {
		fmt.Println("ERROR: Cannot interpret SPEW_INTERVAL " + err.Error())
	}

	// loop
	for {
		fmt.Println(time.Now().Format(time.RFC3339) + " " + spewOut)
		time.Sleep(spewTimespan)
	}

}
