package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"villagertournamentbot/client"
)

func registerWebhook() {
	fmt.Println("Registering webhook...")
	httpClient := client.Create()

	//Set parameters
	path := "https://api.twitter.com/1.1/account_activity/all/" + os.Getenv("WEBHOOK_ENV") + "/webhooks.json"
	values := url.Values{}
	values.Set("url", os.Getenv("APP_URL")+"/webhook/twitter")

	//Make Oauth Post with parameters
	resp, _ := httpClient.PostForm(path, values)
	defer resp.Body.Close()
	//Parse response and check response
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println("Webhook id of " + data["id"].(string) + " has been registered")
	subscribeWebhook()
}
