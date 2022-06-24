package twitter

import (
	"fmt"
	"net/url"
)

func FetchTweetsByAccount(account string) {
	values := url.Values{}
	values.Set("screen_name", account)

	tweets, err := twitterApi.GetUserTimeline(values)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		fmt.Printf("\x1b[31m%s\x1b[0m", tweet.FullText)
		fmt.Print("\n\n")
		// fmt.Print(tweet.FullText, "\n\n")
	}
}
