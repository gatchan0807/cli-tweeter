package tweet

import (
	"os"
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/ahaha0807/cli-tweeter/util"
	"github.com/ChimeraCoder/anaconda"
)

// This method needs for a tweet content to Twitter REST API.
// All tweet request are sent to this method.
// If request has account flag then request is forwarded tweetWithAccount method.
func Tweet(context *cli.Context) error {
	if context.Args().Get(0) == "" {
		log.Fatalf("ツイート内容が存在しませんでした。半角スペース区切りでツイートする内容を入力してください。")
		log.Fatalf("Tweet content does not exist. Please enter your tweet separated by spaces.")
		return nil
	}

	selectedAccountIndex := 0
	if context.String("account") != "" {
		selectedAccountIndex = searchAccount(context.String("account"))
	}

	userInfoList := util.GetUserInfoList()
	tweetAccount := userInfoList[selectedAccountIndex]

	err := doTweet(tweetAccount, context.Args().Get(0))
	util.Check(err)

	return nil
}

func searchAccount(accountId string) int {
	userInfoList := util.GetUserInfoList()

	selectedUserIndex := 0
	for index, element := range userInfoList {
		if element["userId"] == accountId {
			selectedUserIndex = index
		}
	}
	return selectedUserIndex
}

func doTweet(tweetAccount map[string]string, tweetContents string) error {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CLI_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CLI_CONSUMER_SECRET_KEY"))

	api := anaconda.NewTwitterApi(tweetAccount["accessToken"], tweetAccount["accessSecret"])

	tweet, err := api.PostTweet(tweetContents, nil)
	if err == nil {
		fmt.Print("[Tweet by " + tweetAccount["userId"] + " Successed] ")
		fmt.Println(tweet.Text)
		return nil
	} else {
		return err
	}
}
