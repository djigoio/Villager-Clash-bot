package webhook

import (
	"fmt"
	"io/ioutil"
	"os"
)

func subscribeWebhook() {
	fmt.Println("Subscribing webapp...")
	client := CreateClient()
	path := "https://api.twitter.com/1.1/account_activity/all/" + os.Getenv("WEBHOOK_ENV") + "/subscriptions.json"
	resp, _ := client.PostForm(path, nil)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	//If response code is 204 it was successful
	if resp.StatusCode == 204 {
		fmt.Println("Subscribed successfully")
	} else if resp.StatusCode != 204 {
		fmt.Println("Could not subscribe the webhook. Response below:")
		fmt.Println(string(body))
	}
}
