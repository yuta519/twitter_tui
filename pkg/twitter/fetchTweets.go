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

func UnionTweets(array1 []Tweet, array2 []Tweet) []Tweet {
	s := make(map[Tweet]struct{}, len(array1))
	for _, data := range array1 {
		s[data] = struct{}{}
	}
	for _, data := range array2 {
		s[data] = struct{}{}
	}

	newArray := make([]Tweet, 0, len(s))
	for key, _ := range s {
		newArray = append(newArray, key)
	}
	return newArray
}
