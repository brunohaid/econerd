package econonerd

import (
    "fmt"
    "log"
    "net/http"
    "crawler"
    "time"
)

type Item struct {
	ID 				string			`json:"id"`
	Title			string			`json:"title"`
	Author			[]Person 		`json:"author"`
	FirstSeen		time.Time		`json:"firstseen"`	
	URL				string			`json:"url"`	
	mentions		[]Mention 		`json:"mentions"`
}

type Mention struct {
	Person 			Person			`json:"by"`
	URL     		string			`json:"url"`
	Timestamp		time.Time 		`json:"timestamp"`
}

type Person struct {
	Handle 			string			`json:"handle"`
	Name 			string			`json:"name"`
	Aliases			[]string		`json:"aka"`
}

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