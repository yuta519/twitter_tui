package twitter

import (
	"net/url"
)

func FetchTweetsByAccount(account string) []Tweet {
	var tweetsByAccount []Tweet

	values := url.Values{}
	values.Set("screen_name", account)

	tweets, err := twitterApi.GetUserTimeline(values)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		tweetsByAccount = append(
			tweetsByAccount,
			Tweet{
				Id:        tweet.IdStr,
				CreatedAt: tweet.CreatedAt,
				UserName:  tweet.User.Name,
				TweetText: tweet.FullText,
			},
		)
	}
	return tweetsByAccount
}

func FetchHomeTweets() []Tweet {
	var homeTweets []Tweet

	tweets, err := twitterApi.GetHomeTimeline(url.Values{})
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
