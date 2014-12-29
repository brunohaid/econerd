package crawler

import (
	// Core libs
	"time"
	"log"

	// The Interwebs   
	rss "github.com/jteeuwen/go-pkg-rss"
)

// List blogs 
// TODO: Move into config file / DB
var (
	blogs = [...]string{
		"http://alephblog.com/feed/",
		"http://baselinescenario.com/feed/",
		"http://blog.mpettis.com/feed/",
		"http://blogs.reuters.com/anatole-kaletsky/feed/",
		"http://blogs.wsj.com/economics/feed/",
		"http://coppolacomment.blogspot.com/feeds/posts/default?alt=rss",
		"http://crookedtimber.org/feed/",
		"http://delong.typepad.com/sdj/atom.xml",
		"http://equitablegrowth.org/feed/",
		"http://fistfulofeuros.net/feed/",
		"http://ftalphaville.ft.com/feed/",
		"http://johnquiggin.com/feed/",
		"http://krugman.blogs.nytimes.com/feed/",
		"http://maxspeak.net/feed/",
		"http://thereformedbroker.com/feed/",
		"http://yanisvaroufakis.eu/feed/",

		"http://feeds.feedburner.com/BronteCapital",
		"http://feeds.feedburner.com/EconomistsView",		
		"http://feeds.feedburner.com/dealbreaker",
		"http://feeds.feedburner.com/espeak",
		"http://feeds.feedburner.com/JaredBernstein",
		"http://feeds.feedburner.com/MacroAndOtherMarketMusings",
		"http://feeds.feedburner.com/marginalrevolution",
		"http://feeds.feedburner.com/NakedCapitalism",
		"http://feeds.feedburner.com/neweconomicperspectives/yMfv",
		"http://feeds.feedburner.com/TheBigPicture",				

		"http://www.bloombergview.com/rss",		
		"http://www.economist.com/blogs/freeexchange/atom.xml",
		"http://www.economist.com/sections/economics/rss.xml",
		"http://www.huffingtonpost.com/author/index.php?author=ben-walsh",
		"http://www.interfluidity.com/feed",
		"http://www.nextnewdeal.net/rss.xml",
		"https://www.pehub.com/feed/",
		"http://www.project-syndicate.org/rss",
		"http://www.slate.com/all.fulltext.matthew_yglesias.rss",
		"http://www.vox.com/authors/matthew-yglesias/rss",
	}
)


// Crawl a single blog from the list, look for optional digest items
func crawlblogs() {
	//go spawnblogsubscriber("http://coppolacomment.blogspot.com/feeds/posts/default?alt=rss")
	for _, blog := range blogs {
		go spawnblogsubscriber(blog)
	}	
}

// Spawns a new reader
func spawnblogsubscriber(uri string) {
	// Spawn a new feed
	feed := rss.New(10, true, bloghandler, blogposthandler);

	// Log start
	log.Printf("Spawning subscriber to %s",uri)	

	// Loop endlessly
	for {
		// Fetch feed and log errors
		if err := feed.FetchClient(uri, gethttpclient(), nil); err != nil {
			log.Printf("Error fetching %s: %+v", uri, err)
		}		

		// Wait until next update
		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}	
}

// Changes to bloags
func bloghandler(feed *rss.Feed, newchannels []*rss.Channel) {
	log.Printf("Subscribed to %d new channel(s) in %s", len(newchannels), feed.Url)
}

// Handling new items
func blogposthandler(feed *rss.Feed, ch *rss.Channel, newposts []*rss.Item) {
	var content string

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
			author:		post.Author.Name,
			published: 	TimeFromString(post.PubDate),
			title:		post.Title,			
			body: 		content,
		}

		// If we have multiple links
		if len(post.Links) > 1 {
			// Cycle through them
			for _, link := range post.Links {
				// And try to assign alternate
				if link.Rel == "alternate" { item.url = link.Href }
			}
		}

		// Always fall back to first link
		if item.url == "" {
			item.url = post.Links[0].Href
		}

		post := Post{
			item: item,
		}

		// Send it off for processing
		go AddPost(post)
	}
}