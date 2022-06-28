package twitter

import "fmt"

func PrintDiffTweets(newTweets []Tweet, oldTweets []Tweet) {
	slice := make(map[Tweet]struct{}, len(newTweets))
	for _, data := range oldTweets {
		slice[data] = struct{}{}
	}

	result := make([]Tweet, 0, len(oldTweets))
	for _, data := range newTweets {
		if _, ok := slice[data]; ok {
			continue
		}
		result = append(result, data)
	}

	for _, tweet := range result {
		fmt.Println(tweet.UserName)
		fmt.Println(tweet.CreatedAt)
		fmt.Println(tweet.TweetText)
		fmt.Print("\n")
	}
}
