package main

import (
	"fmt"
	"time"

	"github.com/yuta519/twitter_tui/pkg/twitter"
)

func main() {
	tweets := twitter.FetchHomeTweetsByAccount()
	fmt.Println(tweets)
	for range time.Tick(1 * time.Minute) {
		fmt.Println(tweets)
		twitter.FetchHomeTweetsByAccount()
	}
}
