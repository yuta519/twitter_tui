package main

import (
	"fmt"
	"log"

	"github.com/yuta519/twitter_tui/pkg/twitter"
)

func main() {
	var account string
	fmt.Printf("Please input twitter account: ")
	_, err := fmt.Scanln(&account)

	if err != nil {
		log.Fatal(err)
	}

	twitter.PrintTweets(twitter.FetchTweetsByAccount(account))
}
