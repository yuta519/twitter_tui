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
				Id:        utils.ColoredText(tweet.IdStr, onesPlace),
				CreatedAt: utils.ColoredText(tweet.CreatedAt, onesPlace),
				UserName:  utils.ColoredText(tweet.User.Name, onesPlace),
				TweetText: utils.ColoredText(tweet.FullText, onesPlace),
			},
		)
	}

	return timelines
}
