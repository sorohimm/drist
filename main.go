package main

import (
	"errors"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"regexp"
	"time"
)

const telegramTokenEnv string = "5075773271:AAGoTk9skey89rzVoCVz-L6W3EISm0donLs"

func modFilenameForList(fname string) string {
	if fname[len(fname)-4:] == ".jpg" {
		return "PIC " + fname[:len(fname)-4]
	} else {
		return "GIF " + fname[:len(fname)-4]
	}
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		URL: "https://api.telegram.org",

		Token:  telegramTokenEnv,
		Poller: &tb.LongPoller{Timeout: 10 * time.Microsecond},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/drist_list", func(m *tb.Message) {
		file, err := os.Open("./memes")
		if err != nil {
			log.Fatalf("failed opening directory: %s", err)
		}
		defer file.Close()

		list, _ := file.Readdirnames(0) // 0 to read all files and folders

		var dlist string
		for _, f := range list {

			dlist += modFilenameForList(f)
			dlist += "\n"
		}
		b.Send(m.Chat, dlist)
	})

	b.Handle("/drist", func(m *tb.Message) {
		if m.Text == "/drist" {
			a := &tb.Photo{File: tb.FromDisk(fmt.Sprintf("./memes/Drist.jpg"))}
			b.Send(m.Chat, a)
			return
		}

		matched, _ := regexp.MatchString(`^/drist\s[a-z]{1,10}$`, m.Text)
		if !matched {
			b.Send(m.Chat, "Ты че прислал, еблан?")
			return
		}

		name := GetDristName(m.Text)
		if _, err := os.Stat(fmt.Sprintf("./memes/%s.jpg", name)); !errors.Is(err, os.ErrNotExist) {
			a := &tb.Photo{File: tb.FromDisk(fmt.Sprintf("./memes/%s.jpg", name))}
			b.Send(m.Chat, a)
			return
		}

		if _, err := os.Stat(fmt.Sprintf("./memes/%s.gif", name)); !errors.Is(err, os.ErrNotExist) {
			a := &tb.Video{File: tb.FromDisk(fmt.Sprintf("./memes/%s.gif", name))}
			b.Send(m.Chat, a)
			return
		}

		b.Send(m.Chat, "Такого дриста пока нет")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if !m.IsReply() {
			return
		}

		matched, _ := regexp.MatchString(`^#drist\s[a-z]{1,10}$`, m.Text)
		if !matched {
			return
		}

		if m.ReplyTo.Photo != nil {
			NewPhotoDrist(b, m)
		} else if m.ReplyTo.Animation != nil {
			NewAnimDrist(b, m)
		} else {
			return
		}
	})

	b.Start()
}
