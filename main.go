package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sclevine/agouti"
)

type LineConf struct {
	ChannelSecret      string `json:"channelSecret"`
	ChannelAccessToken string `json:"channelToken"`
	UserID             string `json:"userid"`
}

func main() {
	// chrome起動
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--window-size=120,100", // ウィンドウサイズの指定
			"--window-position=-2000000,-200000",
		}),
	)

	if err := driver.Start(); err != nil {
		log.Fatalln(err)
		return
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for {
		if err := page.Navigate("https://reserve.tokyodisneyresort.jp/ticket/search/"); err != nil {
			log.Fatalln(err)
			return
		}

		count, err := page.FindByID("searchTab").Count()
		log.Println(count, err)

		if count >= 1 {
			defer notify()
			break
		}

		if err := page.Reset(); err != nil {
			log.Fatalln(err)
			return
		}

		//頻繁にアクセスするとAccessDeniedになってしまうため
		time.Sleep(time.Millisecond * 400)
	}

}

func notify() {
	l := LineConf{}
	jsonString, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonString, &l)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(l.ChannelSecret, l.ChannelAccessToken, l.UserID)

	line, err := linebot.New(l.ChannelSecret, l.ChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	postMessage := linebot.NewTextMessage("Opened Screen. Check PC.")
	_, err = line.PushMessage(l.UserID, postMessage).Do()
	if err != nil {
		log.Println("post:", err)
	}
}
