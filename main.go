package main

import (
	"fmt"
	"log"
	"net/http"
	"crawler"
)

func init() {
	// Logio
	log.Printf("Econonerd awakens")

	// Init crawler
	crawler.Init()

	// Route calls (also keeps the server alive indefinitely)
	http.HandleFunc("/firehose", firehose)    
	http.HandleFunc("/curated", curated) 
}

func firehose(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm the firehose, weeeeeeeeeeeh")
}

func curated(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Curation. I should do curation.")
}