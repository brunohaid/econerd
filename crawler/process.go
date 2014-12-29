package crawler

import (
	"log"
	"strings"

	// To create unique IDs from URLs
	"crypto/md5"
	"encoding/hex"
)

var (
	// Cache processed IDs
	knownids 	map[string]bool

	// Cache raw URLs already processed
	urltoid		map[string]string	

	// URL endings we remove
	trimmings = [...]string{ 
		// Analytics
		"utm_",
		// Reuters source
		"feedType=RSS",
		// NYtimes sharing
		"hp&amp;action",	
		// WSJ stuff 
		"mod=WSJBlog",
		// Economist RSS tracker
		"fsrc=rss",
	}

	// URLs we do not resolve if they conatin one of those strings 
	// eg because they are redirecting to paywall etc
	donotresolve = [...]string{
		// FTAV
		"ftalphaville.ft.com",
	}	
)

// If a new item was found
func AddPost(post Post) {
	// Resolve final destination trimmed URL
	post.item.resolve()

	// Generate id if none exists yet
	if post.item.id == "" { post.item.hash() }
	
	// Log it
	log.Printf("ITEM on %#v, %#v",post.item.id, post.item.url)
}

// If a new mention was found
func AddMention(mention Mention) {
	// Log it
	log.Printf("MENTION on %s: %s",mention.item.kind,trimurl(mention.target))
}


// Change the URL of an item to it's trimmed final destination
func (i *Item) resolve() {
	// First check if a do not resolve rule applies to raw url
	for _, dnr := range donotresolve {
		// If the dnr string is part of our url
		if strings.Contains(i.url, dnr) {
			return
		}
	}

	// Request URL to find out if there are any redirects
	response, _ := fetchurl(i.url)

	// Extract URL object from response (followed 302)
	url := response.Request.URL

	// Set the items url field to trimmed URL (removing request params etc)
	i.url = trimurl(url.String())	
}

// Change the URL of an item to it's trimmed final destination
func (i *Item) hash() {
	// Fetch URL
	url := strings.ToLower(i.url)

	// Remove  http(s)://, might switch this to regex
	url = strings.Replace(url, "https://", "", 1)
	url = strings.Replace(url, "http://", "", 1)

	// Convert url string to byte slice and then hash
	hash := md5.Sum([]byte(url))

	// Assign ID
	i.id = hex.EncodeToString(hash[:16])
}

// Remove redundant parts from url string
func trimurl(url string) (trimmed string) {
	// Iterate through our trim vars
	for _, pattern := range trimmings {
		// See if we find a a pattern
		if i := strings.Index(url,"?" + pattern); i > -1 {
			url = url[:i]
		}
	}
	return url
}

// Takes URL and tries to reduce it to a canonical core
// This does not have to be a valid URL anymore, just a unique identifier
func GetCanoncical(s string) string {
	return s
}

