package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	twitterConsumerKey := os.Getenv("twitterConsumerKey")
	twitterConsumerSecret := os.Getenv("twitterConsumerSecret")
	twitterAccessToken := os.Getenv("twitterAccessToken")
	twitterAccessSecret := os.Getenv("twitterAccessSecret")

	config := oauth1.NewConfig(twitterConsumerKey, twitterConsumerSecret)
	token := oauth1.NewToken(twitterAccessToken, twitterAccessSecret)

	// base http client
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client based on http client above
	twitterClient := twitter.NewClient(httpClient)

	// your tweet message here
	_, resp, err := twitterClient.Statuses.Update("tchaki tchaki", nil)

	if resp.StatusCode != 200 {
		log.Fatalf("Error on sending tweet")
	}

	log.Print("Your tweet was sent")
}
