package main

import (
	"fmt"
	"time"

	"github.com/yuta519/twitter_tui/pkg/twitter"
)

func main() {
	// var currentTweets []twitter.Tweet
	tweets := twitter.FetchHomeTweets()
	fmt.Println(tweets)
	for range time.Tick(1 * time.Minute) {
		twitter.FetchHomeTweets()
		// updatedTweets = twitter.FetchHomeTweets()
		// hogehoge(updatedTweets, currentTweets)
		// currentTweets = append(currentTweets, updatedTweets)
		// fmt.Println(tweets)
	}
}
