# Install

```bash
go get github.com/ahaha0807/cli-tweeter/cmd/tweeter
```

# Usage

```
tweeter account # アカウント登録 (Account register)
tweeter account -d # アカウント削除 (Delete registered account)

tweeter list # 登録されているアカウントの一覧表示 (Display registered accounts)

tweeter "[Tweet contents]" # ツイートする (Tweet)
tweeter tweet --account [user id] "[Tweet contents]" # アカウント指定でツイートする (Tweet by specified account)
```

# Specifications

tweetコマンド使用時、アカウントを指定がない場合は一番最初に登録されたアカウントでツイートされる。
(If you no specify account then tweet by first registered account.)
