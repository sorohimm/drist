package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var b *tb.Bot

const telegramTokenEnv string = "5091467802:AAFgDlGT-kg95yj_DccVN6g-icsH6FGojHw"

func main() {
	var err error
	b, err = tb.NewBot(tb.Settings{
		URL: "https://api.telegram.org",

		Token:  telegramTokenEnv,
		Poller: &tb.LongPoller{Timeout: 10 * time.Microsecond},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	os.MkdirAll("./memes", os.ModePerm)

	b.Handle("/drist_list", dristlist_handle)
	b.Handle("/drist", drist_handle)
	b.Handle(tb.OnText, generic_handle)
	b.Handle("/yt", yt_handle)

	b.Start()
}
