# disney-reserve




## Overview
コロナになってからdisneyの予約が大分難しくなったので、少しだけ簡略化できれば良いかと思って作りました。  
毎週水曜はアクセス過多でSorryPageに飛ばされてしまいます。正常のページにアクセスできればLineメッセージで伝える簡単なツールです。

### How to start

1. Clone  
```
git clone `this repository`
cd `this repository`
```

2. DriverをInstall + 依存のModuleを導入 
```
brew install chromedriver
go mod tidy
```

3. Lineで送信時に必要なConfigファイルを作成
```
touch configs/config.json
```


4. Line送信時に必要なAccessToken等の設定後に、値をConfigファイルに記載(下記を参照)  
https://www.virment.com/get-accesstoken-line-setup/
```
{
    "channelSecret":"XXXXXXX",
    "channelToken":"XXXXXXXX",
    "userid":"XXXXXXX"
}
```

5. Start

```
go run main.go
```