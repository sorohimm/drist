package main

import (
	"errors"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"regexp"
	"time"
  "path/filepath"
)

const telegramTokenEnv string = "5091467802:AAFgDlGT-kg95yj_DccVN6g-icsH6FGojHw"

func GetDristTypeFromFn(fname string) DristType {
  ext := fname[len(fname) - 4:]
  if ext == ".jpg" {
    return Photo
  } else if ext == ".gif" {
    return Animation
  } else if ext == ".mp4" {
    return Video
  } else {
    return None
  }
}

func modFilenameForList(fname string) string {
  drtype := GetDristTypeFromFn(fname)
  drname := fname[:len(fname) - 4]
	switch drtype {
    case Photo: return "PIC " + drname
    case Animation: return "GIF " + drname
    case Video: return "VID " + drname
    default: return "WTF " + drname
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

	os.MkdirAll("./memes", os.ModePerm)

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

		names, err := filepath.Glob("./memes/" + GetDristName(m.Text, None) + "*")
    if err != nil {
      b.Send(m.Chat, "Дрист не найден, унитаз пуст!")
      return
    }
    if len(names) != 1 {
      b.Send(m.Chat, "Бот захлебнулся, надо вилкой чистить")
    }

		if _, err := os.Stat(names[0]); !errors.Is(err, os.ErrNotExist) {
      fmt.Println(names[0])
      drtype := GetDristTypeFromFn(names[0])
      var drist interface{tb.Sendable}
      switch drtype {
        case Photo: drist = &tb.Photo{File: tb.FromDisk(names[0])}
        case Video: drist = &tb.Video{File: tb.FromDisk(names[0])}
        case Animation: drist = &tb.Animation{File: tb.FromDisk(names[0])}
        default: {
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
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if !m.IsReply() {
			return
		}

		matched, _ := regexp.MatchString(`^#drist\s[a-z]{1,10}$`, m.Text)
		if !matched {
			return
		}

		NewDrist(b, m)
	})

	b.Start()
}

