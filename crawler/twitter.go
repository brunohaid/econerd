package crawler

import (
	// Core libs
		"log"  
	//"time"

	// Comms
	"github.com/ChimeraCoder/anaconda"
)

// Our twitter list
const twitterlist uint32 = 186333198

var api = &anaconda.TwitterApi{}

// Start new twitter crawler
func crawltwitter() {
	// Build the api
	

	// Kick off fetch
	go fetchtwitterlist()	
}

// Fetch the latest tweets
func fetchtwitterlist() {
	searchResult, err := api.GetSearch("golang", nil)
	log.Printf("Results: %#v",searchResult)	
	log.Printf("Error: %#v",err)		
	for _ , tweet := range searchResult {
		log.Println(tweet.Text)
	}	
}

// Try translating an author into a twitter handle
func GetHandle(author string) string {
	return author
}
