package crawler

import (
	// Core
	"time"
	"log"

	// Stream processing
	"io/ioutil"
	"encoding/json"

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

// Get the contents of a specific page
func gethtml( url string ) string {
	if response, err := fetchurl(url); err != nil {
		// If fetching failed
		log.Printf("Error fetching url %s: %+v", url, err)

		// Return empty if something failed
    	return ""
	} else {
		// Always close the body after we read it
        defer response.Body.Close()

        // Read the actual contents, ignore errors for now
        body, _ := ioutil.ReadAll(response.Body)
        
        // Return the contents, converting byte array to string
        return string(body)
	}	
}

// Fetch a JSON url and return it as generic interface
func getjson(url string) ( i interface{}, err error ) {
	// Fetch the URL
	response, err := fetchurl(url); 

	// If we didn't get any data
	if err != nil {
		// Log an error
		log.Printf("Error fetching JSON %s: %+v", url, err)
	} else {
		// Decode bytes b into interface i, override nil error if something goes wrong
		err = json.NewDecoder(response.Body).Decode(&i)
	}

	// Return the interface
	return i, err
}

// Takes URL and tries to reduce it to a canonical core
// This does not have to be a valid URL anymore, just a unique identifier
func GetCanoncical(s string) string {
	return s
}