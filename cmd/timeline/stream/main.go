package main

import (
	"time"

	"github.com/yuta519/twitter_tui/pkg/twitter"
)

func main() {
	var currentTweets []twitter.Tweet
	twitter.PrintDiffTweets(twitter.FetchHomeTweets(), currentTweets)
	currentTweets = twitter.UnionTweets(currentTweets, twitter.FetchHomeTweets())
	for range time.Tick(1 * time.Minute) {
		updatedTweets := twitter.FetchHomeTweets()
		twitter.PrintDiffTweets(updatedTweets, currentTweets)
		currentTweets = twitter.UnionTweets(currentTweets, updatedTweets)
	}
}
