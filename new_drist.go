package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
  "io"
)

type DristType int

const (
  Photo DristType = iota + 1
  Animation
  Video
  None
  Error
)

func GetDristName(text string, drtype DristType) string {
  var fileext string
  switch drtype {
    case Photo: fileext = ".jpg"
    case Animation: fileext = ".gif"
    case Video: fileext = ".mp4"
    case None: fileext = ""
  }
	return text[7:] + fileext
}

func GetDristRc (b *tb.Bot, m *tb.Message) (io.ReadCloser, DristType) {
  var rc io.ReadCloser
  var err error
  var drtype DristType

	if m.Photo != nil {
    rc, err = b.GetFile(m.Photo.MediaFile())
    drtype = Photo
	} else if m.Animation != nil {
		rc, err = b.GetFile(m.Animation.MediaFile())
    drtype = Animation
	} else if m.Video != nil {
    rc, err = b.GetFile(m.Video.MediaFile())
    drtype = Video
  } else {
		return nil, Error
	}

  if err != nil {
    return nil, Error
  }
  return rc, drtype
}

func NewDrist (b *tb.Bot, m *tb.Message) {
  rc, drtype := GetDristRc(b, m.ReplyTo)
  if rc == nil {
		b.Send(m.Chat, "Что-то пошло не так, бот в пиве")
    return
  }

  name := GetDristName(m.Text, drtype)
  if _, err := os.Stat("./memes/" + name); err == nil {
    b.Send(m.Chat, "Такой дрист уже существует.")
    return
  }
  err := SaveDrist(rc, name)
  if err != nil {
    b.Send(m.Chat, "Бот в пиве, мать в канаве")
    return
  }

  b.Send(m.Chat, fmt.Sprintf("Новый дрист %s добавлен", name[:len(name) - 4]))
}

