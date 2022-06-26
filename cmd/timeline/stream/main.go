package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yuta519/twitter_tui/pkg/twitter"
)

func main() {
	var account string
	fmt.Printf("Please input twitter account: ")
	_, err := fmt.Scanln(&account)

	if err != nil {
		log.Fatal(err)
	}
	twitter.FetchHomeTweetsByAccount(account)
	for range time.Tick(1 * time.Minute) {
		twitter.FetchHomeTweetsByAccount(account)
	}
}
