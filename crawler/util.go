package crawler

import "time"

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

		// Use current date as fallback
		return time.Now()
}