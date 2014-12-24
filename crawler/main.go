package crawler

import (
	"log"
	"time"
) 

type Item struct {
	id 				string			`json:"id"`
	title			string			`json:"title"`
	author			[]Person 		`json:"author"`
	published		time.Time		`json:"published"`
	firstseen		time.Time		`json:"firstseen"`	
	url				string			`json:"url"`	
	mentions		[]Mention 		`json:"mentions"`
	body			string			`json:"body"`
}

type Mention struct {
	Person 			Person			`json:"by"`
	URL     		string			`json:"url"`
	Timestamp		time.Time 		`json:"timestamp"`
}

type Person struct {
	handle 			string			`json:"handle"`
	name 			string			`json:"name"`
	aliases			[]string		`json:"aka"`
}

// Settings definition
type Config struct {
    blogs    []string
    twitterlist   string
}

// Init
func Init() {
	// Log start
	log.Println("Spinning up crawler")

	// Kick off twitter crawling
	Crawltwitter()

	// Spawn RSS reader routines
	Crawlblogs()			
}