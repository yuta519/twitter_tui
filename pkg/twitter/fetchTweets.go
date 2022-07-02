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

func FetchHomeTweets() []Tweet {
	var homeTweets []Tweet
	values := url.Values{}

	tweets, err := twitterApi.GetHomeTimeline(values)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		homeTweets = append(
			homeTweets,
			Tweet{
				Id:        tweet.IdStr,
				CreatedAt: tweet.CreatedAt,
				UserName:  tweet.User.Name,
				TweetText: tweet.FullText,
			},
		)
	}
	return homeTweets
}

func UnionTweets(tweets1 []Tweet, tweets2 []Tweet) []Tweet {
	s := make(map[string]Tweet, len(tweets1))
	for _, data := range tweets1 {
		s[data.Id] = data
	}
	for _, data := range tweets2 {
		s[data.Id] = data
	}

	newTweets := make([]Tweet, 0, len(s))
	for _, value := range s {
		newTweets = append(newTweets, value)
	}
	return newTweets
}
