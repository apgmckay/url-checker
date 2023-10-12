package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: url_checker <url1> <url2> <url3> ...")
		os.Exit(1)
	}

	logFile, err := os.OpenFile("url_check.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			logger.Printf("Error checking %s: %s\n", url, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			logger.Printf("%s is up and running (Status Code: %d)\n", url, resp.StatusCode)
		} else {
			logger.Printf("%s is not up (Status Code: %d)\n", url, resp.StatusCode)
		}
	}
}
