package crawler

import (
	"log"
	"time"
) 

// We keep them super simple for now
type Item struct {
	id				string			`json:"id"`
	kind			string			`json:"type"`
	title			string			`json:"title"`
	author			string			`json:"author"`
	published		time.Time		`json:"published"`
	firstseen		time.Time		`json:"firstseen"`
	url				string			`json:"url"`
	body			string			`json:"body"`
}

type Mention struct {
	id				string			`json:"id"`	
	itemid			string			`json:"itemid"`
	kind			string			`json:"type"`
	author			string			`json:"by"`
	url				string			`json:"url"`
	timestamp		time.Time 		`json:"timestamp"`
	body			string			`json:"body"`
}

type Person struct {
	twitter			string			`json:"twitter"`
	reddit			string			`json:"reddit"`
	name 			string			`json:"name"`
	bio				string			`json:"bio"`
	urls			[]string 		`json:"urls"`	
}

type Alias struct {
	name 			string			`json:"name"`
	handle 			string			`json:"handle"`	
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