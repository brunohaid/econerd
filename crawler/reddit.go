package crawler

import (
	// Core libs
	"time"
	"log"

	// Unwrap
	"encoding/json"
)

// Define
type subreddit struct {
	Name 		string
	Latest		float64

}	

var (
	// The subreddits we want to watch
	subreddits = [2]string{ "economics", "finance" }
)

const (
	redditbaseurl string = "https://www.reddit.com"
	redditchannel string = "/new/.json"
)	

// Init reddit crawler
func crawlreddit() {
	// Spawn a subscriber for each subreddit
	for _, r := range subreddits {

		// New subredit struct
		sr := subreddit{
			Name:		r,
		}

		// Create routine
		go spawnredditor(sr)
	}	
}

// Spawns a new redditor
func spawnredditor(sr subreddit) {	
	// Log start
	log.Printf("Started listening to subreddit %s",sr.Name)

	// Unwrap JSON
	type structure struct {
	    Data struct {
	    	Children []struct {
	    		Data struct {
	    			Id 			string
	    			Self 		bool		`json:"is_self"`	
	    			Url 		string
	    			Ts			float64		`json:"created_utc"`
	    			Author 		string
	    			Title 		string
	    			Selftext	string
	    			Permalink	string
	    		}
	    	}
	    }
	}	

	// Define URI
	uri := redditbaseurl + "/r/" + sr.Name + redditchannel

	// Loop endlessly
	for {	
		// Fetch JSON
		raw, _ := fetchurl(uri)		    		

		// Instantiate new structure
		var s structure

		// Decode JSON
		json.NewDecoder(raw.Body).Decode(&s)					

		// Iterate through items
		for _, child := range s.Data.Children {
			// See if it's newer than the last we knew
			if child.Data.Ts <= sr.Latest { continue }

			// Build proper item, map fields
			item := Item{
				id:			child.Data.Id,
				kind:		"reddit",		
				author:		child.Data.Author,
				published: 	time.Unix(int64(child.Data.Ts),0),
				title:		child.Data.Title,					
				body: 		child.Data.Selftext,
			}

			// If it's a posting instead of a link, start working on creating an Item				
			if child.Data.Self {
				// Set Post URL
				item.url = child.Data.Url			

				// Build post
				post := Post{
					item: item,
				}		

				// Send it off for processing
				go AddPost(post)				

			} else {
				// Set mention post url
				item.url = redditbaseurl + child.Data.Permalink

				// Build proper Mention
				mention := Mention{
					item: item,
					target: child.Data.Url,			
				}		

				// Send it off for processing
				go AddMention(mention)				
			}
		} 

		// Set the timestamp to the newest we processed
		sr.Latest = s.Data.Children[0].Data.Ts

		// Wait until next update
		time.Sleep(time.Minute)
	}	
}
