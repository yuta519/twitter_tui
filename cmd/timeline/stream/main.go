package main

import (
	"time"

	"github.com/yuta519/twitter_tui/pkg/twitter"
)

func main() {
	var currentTweets []twitter.Tweet
	updatedTweets := twitter.FetchHomeTweets()
	twitter.PrintDiffTweets(updatedTweets, currentTweets)
	for range time.Tick(1 * time.Minute) {
		updatedTweets := twitter.FetchHomeTweets()
		twitter.PrintDiffTweets(updatedTweets, currentTweets)
		// Need to create union function
		//currentTweets = append(currentTweets, updatedTweets)
	}
}
