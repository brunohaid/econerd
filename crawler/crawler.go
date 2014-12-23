package crawler

import (
    "fmt"
    "time"
    "net/http"
    "log"
)

func Crawl(w http.ResponseWriter, r *http.Request) {
	// Get current timestamp
	t := time.Now().UnixNano()
	// Console log
	log.Println("Starting new crawling session at ")
	// Cofirmation page
    fmt.Fprint(w, "Off we go, let's crawl this shit \nNew crawling session started at ",t)
}