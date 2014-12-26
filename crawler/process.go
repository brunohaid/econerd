package crawler

import "log"

// If a new item was found
func AddItem(item Item) {
	// Log it
	log.Printf("ITEM on %s: %s",item.kind,item.url)	
}

// If a new mention was found
func AddMention(mention Mention) {
	// Log it
	log.Printf("MENTION on %s: %s",mention.kind,mention.target)	
}

