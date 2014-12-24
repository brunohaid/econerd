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
	go spawnsubscriber("http://www.interfluidity.com/feed")
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
func itemhandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	log.Printf("%d new item(s) in %s", len(newitems), feed.Url)
    for _, item := range newitems {
    	url := item.Links[0].Href
        log.Println(item.Title,item.Author.Name,item.PubDate,url)
    }
}