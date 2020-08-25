**※ 2020/08/25追記 現在、ベースに使っているCLIライブラリのアップデートがあった影響で動かなくなっています。追ってアップデートし、配信します**
https://github.com/urfave/cli

# Install

```bash
go get github.com/gatchan0807/cli-tweeter/cmd/tweeter
```

# Usage

```
tweeter account # アカウント登録 (Account register)
tweeter account -d # アカウント削除 (Delete registered account)

tweeter list # 登録されているアカウントの一覧表示 (Display registered accounts)

tweeter "[Tweet contents]" # ツイートする (Tweet)
tweeter tweet --account [user id] "[Tweet contents]" # アカウント指定でツイートする (Tweet by specified account)
```

# Preparation
（サーバー作って公開してそこにアクセスしてもらうやり方にしても良かったのですが、めんどくさかったので、）

このツールを使用するには**Twitter APIキーを取得し、環境変数に登録する必要があります。**
やり方は[こちらのサイト](http://phiary.me/twitter-api-key-get-how-to/)を参考にAPIキーを取得し、
以下の2つのコマンドで取得したTwitter APIキーを登録しておいてください。

You need a **Consumer key(API key)** and **Consumer secret(API secret)** of Twitter for using this tool.
Before use this tool, Please get Consumer key(API key) and Consumer secret(API secret) and register to your environment variables. 

```
export TWITTER_CLI_CONSUMER_KEY=[Consumer Key(API Key)] >> ~/.bash_profile # or .your_profile
export TWITTER_CLI_CONSUMER_SECRET_KEY=[Consumer Secret(API Secret)] >> ~/.bash_profile # or .your_profile
```

# Specifications

tweetコマンド使用時、アカウントを指定がない場合は一番最初に登録されたアカウントでツイートされます。
(If you no specify account then tweet by first registered account.)
