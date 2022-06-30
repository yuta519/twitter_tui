package twitter

import (
	"fmt"

	"github.com/yuta519/twitter_tui/pkg/utils"
)

func PrintDiffTweets(newTweets []Tweet, oldTweets []Tweet) {
	comparisonMap := make(map[string]Tweet, len(newTweets))
	for _, data := range oldTweets {
		comparisonMap[data.Id] = data
	}

	tweets := make([]Tweet, 0, len(oldTweets))
	for _, data := range newTweets {
		if _, ok := comparisonMap[data.Id]; ok {
			continue
		}
		tweets = append(tweets, data)
	}

	for i := len(tweets) - 1; i >= 0; i-- {
		onesPlace := i % 10
		if i%10 >= 8 {
			onesPlace = 10 - i%10
		}

		fmt.Println(utils.ColoredText(tweets[i].CreatedAt, onesPlace))
		fmt.Println(utils.ColoredText(tweets[i].UserName, onesPlace))
		fmt.Println(utils.ColoredText(tweets[i].TweetText, onesPlace))
		fmt.Print("\n")
	}
}
