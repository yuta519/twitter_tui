package twitter

import (
	"fmt"
	"net/url"
	"strconv"
)

func FetchTweetsByAccount(account string) {
	values := url.Values{}
	values.Set("screen_name", account)

	tweets, err := twitterApi.GetUserTimeline(values)
	if err != nil {
		panic(err)
	}

	for i, tweet := range tweets {
		onesPlace := strconv.Itoa(i % 10)
		coloredTweet := fmt.Sprintf(
			"%s%s%s%s%s", "\x1b[3", onesPlace, "m", tweet.FullText, "\x1b[0m",
		)
		fmt.Printf(coloredTweet)
		fmt.Print("\n\n")
	}
}
