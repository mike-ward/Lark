package models

import (
	twitter2 "Lark/twitter"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"time"
)

var HomeTweets []twitter.Tweet
var LikeTweets []twitter.Tweet

func TimelineFetch() {
	userTweets()
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		userTweets()
	}
}

func userTweets() {
	isAuthenticated, _ := LarkSettings.IsAuthenticated.Get()

	if isAuthenticated {
		config := oauth1.NewConfig(twitter2.ConsumerKey, twitter2.ConsumerSecret)
		token := oauth1.NewToken(LarkSettings.AccessToken, LarkSettings.AccessTokenSecret)
		httpClient := config.Client(oauth1.NoContext, token)
		client := twitter.NewClient(httpClient)
		tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{Count: 20})

		if err == nil {
			HomeTweets = tweets
		} else {
			fmt.Println(resp)
		}
	}
}
