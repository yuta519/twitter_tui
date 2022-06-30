package twitter

import "fmt"

func PrintDiffTweets(newTweets []Tweet, oldTweets []Tweet) {
	slice := make(map[Tweet]struct{}, len(newTweets))
	for _, data := range oldTweets {
		slice[data] = struct{}{}
	}

	tweets := make([]Tweet, 0, len(oldTweets))
	for _, data := range newTweets {
		if _, ok := slice[data]; ok {
			continue
		}
		tweets = append(tweets, data)
	}

	for i := len(tweets) - 1; i >= 0; i-- {
		fmt.Println(tweets[i].UserName)
		fmt.Println(tweets[i].CreatedAt)
		fmt.Println(tweets[i].TweetText)
		fmt.Print("\n")
	}
	// for _, tweet := range tweets {
	// 	fmt.Println(tweet.UserName)
	// 	fmt.Println(tweet.CreatedAt)
	// 	fmt.Println(tweet.TweetText)
	// 	fmt.Print("\n")
	// }
}
