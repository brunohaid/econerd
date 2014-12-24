package crawler

import "log"

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
	// Crawlblogs()			
}