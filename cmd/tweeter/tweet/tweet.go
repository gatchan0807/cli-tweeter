package tweet

import (
	"github.com/urfave/cli"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
	"github.com/ChimeraCoder/anaconda"
)

var accountListFilePath = "/tmp/tweeter/user_account.csv"

func Tweet(context *cli.Context) error {

	if context.String("account") != "" {
		TweetWithAccount(context, context.String("account"))
		return nil
	}

	data, err := ioutil.ReadFile(accountListFilePath)
	if err != nil {
		fmt.Println("ユーザーのアカウントは登録されていませんでした。")
		fmt.Println("(User account not found.)")
		return nil
	}

	userInfoList := util.ConvertToInfoList(string(data))
	tweetAccount := userInfoList[0]

	err = doTweet(tweetAccount, context.Args().Get(0))
	util.Check(err)

	return nil
}

func TweetWithAccount(context *cli.Context, accountId string) error {

	fmt.Println("tweet with account")
	fmt.Println(accountId)

	fmt.Println(context.Args().Get(0))

	return nil
}

func doTweet(tweetAccount map[string]string, tweetContents string) error {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CLI_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CLI_CONSUMER_SECRET_KEY"))

	api := anaconda.NewTwitterApi(tweetAccount["accessToken"], tweetAccount["accessSecret"])

	tweet, err := api.PostTweet(tweetContents, nil)
	fmt.Print("[Tweet Successed] ")
	fmt.Println(tweet.Text)

	if err != nil {
		return err
	} else {
		return nil
	}
}
