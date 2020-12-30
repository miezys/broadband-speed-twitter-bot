package pkg

import (
	"broadband-speed-twitter-bot/configs"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func GetTwitterClient(c *configs.Configuration) (*twitter.Client, error) {
	config := oauth1.NewConfig(c.TwitterAPIConfig.ConsumerKey,
		c.TwitterAPIConfig.ConsumerSecret)
	token := oauth1.NewToken(c.TwitterAPIConfig.AccessToken,
		c.TwitterAPIConfig.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}
