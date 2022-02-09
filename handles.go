package main

import (
	"context"
	"errors"
	"fmt"
	yt "github.com/wader/goutubedl"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func drist_handle(m *tb.Message) {
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

	names, err := filepath.Glob("./memes/" + GetDristName(m.Text, None) + "*")
	if err != nil || len(names) < 1 {
		b.Send(m.Chat, "Дрист не найден, унитаз пуст!")
		return
	}

	if _, err := os.Stat(names[0]); !errors.Is(err, os.ErrNotExist) {
		fmt.Println(names[0])
		drtype := GetDristTypeFromFn(names[0])
		var drist interface{ tb.Sendable }
		switch drtype {
		case Photo:
			drist = &tb.Photo{File: tb.FromDisk(names[0])}
		case Video:
			drist = &tb.Video{File: tb.FromDisk(names[0])}
		case Animation:
			drist = &tb.Animation{File: tb.FromDisk(names[0])}
		default:
			{
				b.Send(m.Chat, "Крейзи дрист")
				return
			}
		}
		_, err := b.Send(m.Chat, drist)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// TODO(sorohimm): possibly dead code, see globbing
	b.Send(m.Chat, "Такого дриста пока нет")
}

func dristlist_handle(m *tb.Message) {
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
}

func yt_handle(m *tb.Message) {
	ytdl_ctx, err := yt.New(context.Background(), m.Text[4:len(m.Text)], yt.Options{Type: yt.TypeSingle})
	if err != nil {
		b.Send(m.Chat, "гегель протев, дрисня в пиве")
		return
	}
	dwnld, err := ytdl_ctx.Download(context.Background(), "best")
	if err != nil {
		b.Send(m.Chat, "Бать а как")
		return
	}
	defer dwnld.Close()
	b.Send(m.Chat, &tb.Video{File: tb.FromReader(dwnld)})
}

func generic_handle(m *tb.Message) {
	if !m.IsReply() {
		return
	}

	matched, _ := regexp.MatchString(`^#drist\s[a-z]{1,10}$`, m.Text)
	if !matched {
		return
	}

	NewDrist(b, m)
}
