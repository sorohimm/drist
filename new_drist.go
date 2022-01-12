package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

func NewPhotoDrist (b *tb.Bot, m *tb.Message) {
	rc, err := b.GetFile(m.ReplyTo.Photo.MediaFile())
	if err != nil {
		return
	}

	name := GetNewDristName(m.Text)
	if _, err := os.Stat(fmt.Sprintf("./memes/%s.jpg", name)); err == nil {
		b.Send(m.Chat, "Такой дрист уже существует.")
		return
	}

	err = SavePhotoDrist(rc, name)
	if err != nil {
		b.Send(m.Chat, "Что-то пошло не так, бот в пиве")
		return
	}

	b.Send(m.Chat, fmt.Sprintf("Новый дрист %s добавлен", name))
}

func NewAnimDrist (b *tb.Bot, m *tb.Message) {
	rc, err := b.GetFile(m.ReplyTo.Animation.MediaFile())
	if err != nil {
		return
	}

	name := GetNewDristName(m.Text)
	if _, err := os.Stat(fmt.Sprintf("./memes/%s.gif", name)); err == nil {
		b.Send(m.Chat, "Такой дрист уже существует.")
		return
	}

	err = SaveAnimDrist(rc, name)
	if err != nil {
		b.Send(m.Chat, "Что-то пошло не так, бот в пиве")
		return
	}

	b.Send(m.Chat, fmt.Sprintf("Новый дрист %s добавлен", name))
}
