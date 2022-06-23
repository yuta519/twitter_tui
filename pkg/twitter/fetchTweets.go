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
		fmt.Println("tweet: ", tweet.Text)
	}
}
