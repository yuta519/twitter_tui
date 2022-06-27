package twitter

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var twitterApi *anaconda.TwitterApi

func init() {
	anaconda.SetConsumerKey(os.Getenv("API_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("API_SECRET"))
	twitterApi = anaconda.NewTwitterApi(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"))
}

// func Initialize() *anaconda.TwitterApi {
// 	anaconda.SetConsumerKey(os.Getenv("API_KEY"))
// 	anaconda.SetConsumerSecret(os.Getenv("API_SECRET"))
// 	twitterApi := anaconda.NewTwitterApi(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"))
// 	return twitterApi
// }
