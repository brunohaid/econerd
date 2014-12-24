package crawler

import (
	// Core libs
	"time"
	"log"

	// The Interwebs   
	rss "github.com/jteeuwen/go-pkg-rss"
	"net/http"
)

// Crawl a single blog from the list, look for optional digest items
func Crawlblogs() {
	go spawnsubscriber("http://ftalphaville.ft.com/feed/")
	go spawnsubscriber("http://www.bloombergview.com/rss")
	go spawnsubscriber("https://medium.com/feed/bull-market")
	// go spawnsubscriber("http://www.interfluidity.com/feed")
}

// Spawns a new reader
func spawnsubscriber(uri string) {
	// Spawn a new feed
	feed := rss.New(10, true, feedhandler, itemhandler);

	// Log start
	log.Printf("Spawning subscriber to %s",uri)	

	// Loop endlessly
	for {
		// Build a new client
		transport := http.Transport{}

		client := &http.Client{
			Transport: &transport,
		} 	

		// Fetch feed and log errors
		if err := feed.FetchClient(uri, client, nil); err != nil {
			log.Printf("Error fetching %s: %+v", uri, err)
		}		

		// Wait until next update
		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}	
}

// Changes to bloags
func feedhandler(feed *rss.Feed, newchannels []*rss.Channel) {
	log.Printf("Subscribed to %d new channel(s) in %s", len(newchannels), feed.Url)
}

// Handling new items
func itemhandler(feed *rss.Feed, ch *rss.Channel, newposts []*rss.Item) {
	log.Printf("%d new item(s) in %s", len(newposts), feed.Url)

	// Iterate through new items
	for _, post := range newposts {
		// Read contents from ATOM or fallback on RSS
		var content string

		if post.Content == nil {
			content = post.Description
		} else {
			content = post.Content.Text
		}

		// Get timestamp
		ts, _ := time.Parse("02/Jan/2006:15:04:05 -0700",post.PubDate)

		author := Person{
			name: post.Author.Name,
		}

		// Build proper item
		item := Item{
			url: 		post.Links[0].Href,
			author:		[]Person{author},
			title:		post.Title,
			published: 	ts,
			firstseen:	time.Now().UTC(),
			body: 		content,
		}

		// Send it off for processing
		go Process(item)
	}
}