package crawler

import (
	// Core libs
    "time"
    "log"  

    // Comms
    // "net/http"
)

// Fetch new tweets
func Crawltwitter() {
	// Get current timestamp
	t := time.Now().UnixNano()

	// Log start
	log.Printf("Starting new twitter session at ",t)	
}
