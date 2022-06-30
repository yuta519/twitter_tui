package twitter

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/yuta519/twitter_tui/pkg/utils"
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
	var timelines []Tweet
	values := url.Values{}

	tweets, err := twitterApi.GetHomeTimeline(values)
	if err != nil {
		panic(err)
	}

	for i, tweet := range tweets {
		onesPlace := i % 10
		if i%10 >= 8 {
			onesPlace = 10 - i%10
		}
		timelines = append(
			timelines,
			Tweet{
				Id:        tweet.IdStr,
				CreatedAt: utils.ColoredText(tweet.CreatedAt, onesPlace),
				UserName:  utils.ColoredText(tweet.User.Name, onesPlace),
				TweetText: utils.ColoredText(tweet.FullText, onesPlace),
			},
		)
	}

	return timelines
}

func UnionTweets(tweets1 []Tweet, tweets2 []Tweet) []Tweet {
	s := make(map[Tweet]struct{}, len(tweets1))
	for _, data := range tweets1 {
		s[data] = struct{}{}
	}
	for _, data := range tweets2 {
		s[data] = struct{}{}
	}

	newTweets := make([]Tweet, 0, len(s))
	for key, _ := range s {
		newTweets = append(newTweets, key)
	}
	return newTweets
}
