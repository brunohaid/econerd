package crawler

import (
	// Core
	"time"

	// Interwebs!
	"net/http"
)	

var (
	// The formats we test for
	timeformats = [2]string{ time.RFC1123, time.RFC1123Z }
)

// Takes a string and tries to convert it to time Object
func TimeFromString(s string) (t time.Time) {
		// Iterate through formats defined above
		for _, format := range timeformats {
			var err error
			// Try to parse
			t, err = time.Parse(format,s)
			// If we have no error, return it
			if err == nil {
				return t
			}
		}

		// Use current time as fallback
		return time.Now()
}

// Returns a proper http.client for appengine
func gethttpclient() (client *http.Client) {
	// Build a new client
	transport := http.Transport{}

	client = &http.Client{
		Transport: &transport,
	} 

	return	
}

// A wrapper for http.Get() as we can not use it on appengine 
// but don't want to create a request / response context
func fetchurl(url string) ( *http.Response, error) {
	// First we need a proper client
	client := gethttpclient()

	return client.Get(url)
}