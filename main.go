package main

import (
    "fmt"
    "net/http"
    "crawler"
)

func init() {
    http.HandleFunc("/firehose", firehose)    
    http.HandleFunc("/curated", curated) 
    http.HandleFunc("/crawl", crawler.Crawl)          
}

func firehose(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "I'm the firehose, weeeeeeeeeeeh")
}

func curated(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Curation. I should do curation.")
}