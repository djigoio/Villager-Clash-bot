package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler called")
	//Read the body of the tweet
	body, _ := ioutil.ReadAll(request.Body)
	//Initialize a webhok load obhject for json decoding
	var load WebhookLoad
	err := json.Unmarshal(body, &load)
	if err != nil {
		fmt.Println("An error occured: " + err.Error())
	}
	//Check if it was a tweet_create_event and tweet was in the payload and it was not tweeted by the bot
	if len(load.TweetCreateEvent) < 1 || load.UserId == load.TweetCreateEvent[0].User.IdStr {
		return
	}
	//Send Hello world as a reply to the tweet, replies need to begin with the handles
	//of accounts they are replying to
	_, err = SendTweet("@"+load.TweetCreateEvent[0].User.Handle+" Hello World", load.TweetCreateEvent[0].IdStr)
	if err != nil {
		fmt.Println("An error occured:")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Tweet sent successfully")
	}
}
