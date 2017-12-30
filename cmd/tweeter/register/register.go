package register

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
	"github.com/mrjones/oauth"
	"github.com/skratchdot/open-golang/open"
	"log"
)

func Register(context *cli.Context) error {
	if _, err := os.Stat("/tmp/tweeter/user_account.csv"); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create("/tmp/tweeter/user_account.csv")
		util.Check(err)
	}

	inputUserID := checkUserId()

	if inputUserID == "cancel" {
		return nil
	}

	userAccountToken, userAccountSecret, userID := getTwitterToken()

	if inputUserID != userID {
		fmt.Println("入力されたIDと認証許可したIDが一致しませんでした。")
		fmt.Println("(Didn't match input the user ID and user ID authenticated.)")

		return nil
	}

	success := addToCsvFile(userID, userAccountToken, userAccountSecret)
	if !success {
		fmt.Println("保存失敗しました。")
		fmt.Println("(Save failed)")
	}

	return nil
}

func checkUserId() string {
	var userAccountName string
	fmt.Println("@無しで登録したいTwitterIDを入力してください。")
	fmt.Println("(Input your twitter account ID.(without '@'))")
	fmt.Scan(&userAccountName)

	for userAccountName == "" || isExist(userAccountName) {
		if isExist(userAccountName) {
			fmt.Println(userAccountName + "はすでに登録されています")
			fmt.Println("(" + userAccountName + " is already exist.)")
			fmt.Println("@無しで登録したいTwitterIDを入力してください。登録をキャンセルする場合は ':q' を入力してください。")
			fmt.Println("(Input your Twitter account ID.(without '@') or If you want cancel then type ':q'.)")
			fmt.Scan(&userAccountName)
		}
	}

	if userAccountName == ":q" {
		return "cancel"
	}

	return userAccountName
}

func isExist(userId string) bool {
	data, err := ioutil.ReadFile("/tmp/tweeter/user_account.csv")
	util.Check(err)

	userIdList := util.ConvertToUserIdList(string(data))

	for _, element := range userIdList {
		if element == userId {
			return true
		}
	}
	return false
}

func addToCsvFile(accountName, accountToken, accountSecret string) bool {
	file, err := os.OpenFile("/tmp/tweeter/user_account.csv", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return false
	}

	writeLine := accountName + "," + accountToken + "," + accountSecret + "\n"

	fmt.Fprint(file, writeLine)

	return true
}

func getTwitterToken() (token, secret, userID string) {
	consumer := oauth.NewConsumer(
		os.Getenv("TWITTER_CLI_CONSUMER_KEY"),
		os.Getenv("TWITTER_CLI_CONSUMER_SECRET_KEY"),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	util.Check(err)

	open.Run(url)

	fmt.Println("「連携アプリを認証」ボタンを押した後、ブラウザに表示されたPINコードを入力してください。")
	fmt.Println("(Please push a button for authentication. After input visualized PIN.)")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := consumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		fmt.Println("認証失敗")
		fmt.Println("(Authenticate faild.)")
		log.Fatal(err)
	}

	userID = accessToken.AdditionalData["screen_name"]
	token = accessToken.Token
	secret = accessToken.Secret

	return token, secret, userID
}
