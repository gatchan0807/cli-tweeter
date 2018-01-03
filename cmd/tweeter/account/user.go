package account

import (
	"fmt"
	"os"
	"log"

	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
	"github.com/mrjones/oauth"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
	"github.com/pkg/errors"
)

var accountListFilePath = "/tmp/tweeter/accounts.csv"

// This method is for register Twitter account.
// You call this method, then start account authentication with interpreter.
// And this to open browser for displaying Twitter OAuth PIN number.
// This method create a csv file at `/tmp/tweeter` directory to save account information.
func Account(context *cli.Context) error {
	if context.Bool("delete") {
		err := userDelete()
		util.Check(err)
		return nil
	} else {
		err := userRegister()
		util.Check(err)
		return nil
	}
}

func userRegister() error {
	// check account id file
	if _, err := os.Stat(accountListFilePath); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create(accountListFilePath)
		util.Check(err)
	}

	inputUserId := requestUserId()
	if inputUserId == "cancel" {
		fmt.Println("登録をキャンセルしました。")
		fmt.Println("(Canceled to register.)")
	}

	userAccountToken, userAccountSecret, userID := getTwitterToken()
	if inputUserId != userID {
		fmt.Println("入力されたIDと認証許可したIDが一致しませんでした。")
		fmt.Println("(Didn't match input the account ID and account ID authenticated.)")
	}

	err := util.Push(userID, userAccountToken, userAccountSecret)
	util.Check(err)

	return nil
}

func userDelete() error {
	var userAccountName string
	fmt.Println("削除したいアカウントIDを入力してください")
	fmt.Println("(Please enter account ID you want to delete from the list of accounts.)")
	fmt.Scan(&userAccountName)

	userInfoIndex := util.FindUserIndex(userAccountName)
	if userInfoIndex == -1 {
		return errors.New("Account info not found.")
	}

	userInfoList := util.GetUserInfoList()
	if len(userInfoList) == 0 {
		return nil
	}

	var result []map[string]string
	for index, element := range userInfoList {
		if index != userInfoIndex {
			result = append(result, element)
		}
	}

	err := util.Replace(result)
	util.Check(err)
	if err == nil {
		fmt.Println("削除完了しました。")
		fmt.Println("(Account information has deleted.)")
	}

	return nil
}

func requestUserId() string {
	var userAccountName string
	fmt.Println("登録したいTwitterIDを@無しで入力してください。")
	fmt.Println("(Input your twitter account ID.(without '@'))")
	fmt.Scan(&userAccountName)

	for util.FindUserInfo(userAccountName) != nil {
		fmt.Println(userAccountName + "はすでに登録されています")
		fmt.Println("(" + userAccountName + " is already exist.)")
		fmt.Println("@無しで登録したいTwitterIDを入力してください。登録をキャンセルする場合は ':q' を入力してください。")
		fmt.Println("(Input your Twitter account ID.(without '@') or If you want cancel then type ':q'.)")
		fmt.Scan(&userAccountName)
	}

	if userAccountName == ":q" {
		return "cancel"
	}

	return userAccountName
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
