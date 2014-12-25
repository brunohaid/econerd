package crawler

import "log"

// If a new item was found
func AddItem(item Item) {
	// Log it
	log.Println("ITEM:",item.title)	
}

// If a new mention was found
func AddMention(mention Mention) {
	// Log it
	log.Println("ITEM:",mention)	
}

