package crawler

import (
	"log"
	"strings"
)

var (
	// Cache processed IDs
	knownids map[string]bool

	// URL endings we remove
	trimmings = [...]string{ 
		// Analytics
		"utm_",
		// Reuters source
		"feedType=RSS",
		// NYtimes sharing
		"hp&amp;action" }	
)

// If a new item was found
func AddItem(item Item) {
	// Log it
	log.Printf("ITEM on %s: %s",item.kind,trimurl(item.url))
}

// If a new mention was found
func AddMention(mention Mention) {
	// Log it
	log.Printf("MENTION on %s: %s",mention.kind,trimurl(mention.target))
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

