package twitter

import (
	"fmt"
	"math/rand"

	"github.com/yuta519/twitter_tui/pkg/utils"
)

func PrintTweets(tweets []Tweet) {
	for i := len(tweets) - 1; i >= 0; i-- {
		randint := rand.Intn(8)
		fmt.Println(utils.ColoredText(tweets[i].CreatedAt, randint))
		fmt.Println(utils.ColoredText(tweets[i].UserName, randint))
		fmt.Println(utils.ColoredText(tweets[i].TweetText, randint))
		fmt.Print("\n")
	}
}

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
		randint := rand.Intn(8)
		fmt.Println(utils.ColoredText(tweets[i].CreatedAt, randint))
		fmt.Println(utils.ColoredText(tweets[i].UserName, randint))
		fmt.Println(utils.ColoredText(tweets[i].TweetText, randint))
		fmt.Print("\n")
	}
}
