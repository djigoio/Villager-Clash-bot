package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/dghubble/oauth1"
)

//Struct to parse tweet
type Tweet struct {
	Id    int64
	IdStr string `json:"id_str"`
	Text  string
}

func CreateClient() *http.Client {
	//Create oauth client with consumer keys and access token
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN_KEY"), os.Getenv("ACCESS_TOKEN_SECRET"))

	return config.Client(oauth1.NoContext, token)
}

func SendTweet(tweet string, reply_id string) (*Tweet, error) {
	fmt.Println("Sending tweet as reply to " + reply_id)
	//Initialize tweet object to store response in
	var responseTweet Tweet
	//Add params
	params := url.Values{}
	params.Set("status", tweet)
	params.Set("in_reply_to_status_id", reply_id)
	//Grab client and post
	client := CreateClient()
	resp, err := client.PostForm("https://api.twitter.com/1.1/statuses/update.json", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//Decode response and send out
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &responseTweet)
	if err != nil {
		return nil, err
	}
	return &responseTweet, nil
}
