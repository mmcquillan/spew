package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// init
	fmt.Println("Spewing...")

	// input
	spewPre := "spew"
	if len(os.Args) > 1 {
		spewPre = os.Args[1]
	}

	// force fail
	if os.Getenv("SPEW_FAIL") != "" {
		if strings.ToUpper(os.Getenv("SPEW_FAIL")) == "TRUE" {
			os.Exit(1)
		}
	}

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
		fmt.Println(time.Now().Format(time.RFC3339) + " [" + spewPre + "] " + spewOut)
		time.Sleep(spewTimespan)
	}

}
