package client

import (
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
)

func Create() *http.Client {
	//Create oauth client with consumer keys and access token
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN_KEY"), os.Getenv("ACCESS_TOKEN_SECRET"))

	return config.Client(oauth1.NoContext, token)
}
