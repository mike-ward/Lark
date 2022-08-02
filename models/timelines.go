package models

import (
	larkTwitter "Lark/twitter"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"time"
)

var (
	HomeTweets  []twitter.Tweet
	LikeTweets  []twitter.Tweet
	DataUpdated = make(chan bool)
)

func HomeTimelineLoop() {
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
		HomeTweets = tweets
		DataUpdated <- true
	} else {
		fmt.Println(resp)
	}
}
