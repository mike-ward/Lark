package models

import (
	larkTwitter "Lark/twitter"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"time"
)

var (
	HomeTweets chan []twitter.Tweet
	LikeTweets chan []twitter.Tweet
)

func HomeTimelineLoop() {
	HomeTweets = make(chan []twitter.Tweet)
	getHomeTweets()

	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		if isAuthenticated, _ := LarkSettings.IsAuthenticated.Get(); isAuthenticated {
			getHomeTweets()
		}
	}
}

func getHomeTweets() {
	yes := true
	client := larkTwitter.GetTwitterClient(LarkSettings.AccessToken, LarkSettings.AccessTokenSecret)
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{Count: 75, IncludeEntities: &yes})

	if err == nil {
		HomeTweets <- tweets
	} else {
		fmt.Println(resp)
	}
}
