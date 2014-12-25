package crawler

import (
	// Core libs
	"time"
	"log"

	// The Interwebs   
	rss "github.com/jteeuwen/go-pkg-rss"
	"net/http"
)

var (
	// The subreddits we want to watch
	subreddits = [...]string{ "Economics", "Finance" }
)

// Crawl a single blog from the list, look for optional digest items
func Crawlreddit() {
	// go spawnsubscriber("http://ftalphaville.ft.com/feed/")
	go spawnsubscriber("http://www.bloombergview.com/rss")
	go spawnsubscriber("http://feeds.feedburner.com/EconomistsView")
	// go spawnsubscriber("https://medium.com/feed/bull-market")
	// go spawnsubscriber("http://www.interfluidity.com/feed") 
	// http://thereformedbroker.com/feed/ http://delong.typepad.com/sdj/atom.xml 
	// http://feeds.feedburner.com/TheBigPicture http://feeds.feedburner.com/dealbreaker
	// http://feeds.feedburner.com/marginalrevolution/feed http://alephblog.com/feed/
	// http://feeds.feedburner.com/NakedCapitalism http://www.huffingtonpost.com/author/index.php?author=ben-walsh
	// http://johnquiggin.com/feed/ http://crookedtimber.org/feed/ http://fistfulofeuros.net/feed/ http://feeds.feedburner.com/neweconomicperspectives/yMfv
	// http://yanisvaroufakis.eu/feed/  http://feeds.feedburner.com/espeak http://krugman.blogs.nytimes.com/feed/ http://baselinescenario.com/feed/
	// http://www.vox.com/authors/matthew-yglesias/rss http://www.project-syndicate.org/rss
	// http://www.economist.com/sections/economics/rss.xml http://www.economist.com/blogs/freeexchange/atom.xml 
	// http://blogs.wsj.com/economics/feed/ http://equitablegrowth.org/feed/ http://feeds.feedburner.com/JaredBernstein?format=xml
	// http://feeds.feedburner.com/MacroAndOtherMarketMusings?format=xml http://www.slate.com/all.fulltext.matthew_yglesias.rss
	// http://blogs.reuters.com/anatole-kaletsky/feed/ http://feeds.feedburner.com/blogspot/XqoV http://www.nextnewdeal.net/rss.xml
	// http://maxspeak.net/feed/ https://www.pehub.com/feed/ http://feeds.feedburner.com/BronteCapital?format=xml http://coppolacomment.blogspot.com/feeds/posts/default
	// http://noahpinionblog.blogspot.com/feeds/posts/default http://feeds.feedburner.com/EconomistsView

}

// Spawns a new reader
func spawnsubscriber(subreddit string) {
	// Spawn a new feed
	feed := rss.New(1, true, nil, itemhandler);

	// Log start
	log.Printf("Started listening to subreddit %s",subreddit)	

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

// Handling new items
func itemhandler(feed *rss.Feed, ch *rss.Channel, newposts []*rss.Item) {
	var content string

	// Log
	log.Printf("%d new item(s) in %s", len(newposts), feed.Url)

	// Iterate through new items
	for _, post := range newposts {

		// Read contents from ATOM or fallback on RSS
		if post.Content == nil {
			content = post.Description
		} else {
			content = post.Content.Text
		}

		// Build proper item
		item := Item{
			kind:		"rss",
			url: 		post.Links[0].Href,
			author:		post.Author.Name,
			title:		post.Title,
			published: 	TimeFromString(post.PubDate),
			firstseen:	time.Now(),
			body: 		content,
		}

		// Send it off for processing
		go AddItem(item)
	}
}