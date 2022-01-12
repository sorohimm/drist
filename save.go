package main

import (
	"fmt"
	"io"
	"os"
)

func SavePhotoDrist(rc io.ReadCloser, title string) error {
	file, err := os.Create(fmt.Sprintf("./memes/%s.jpg", title))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, rc)
	return nil
}

func SaveAnimDrist(rc io.ReadCloser, title string) error {
	file, err := os.Create(fmt.Sprintf("./memes/%s.gif", title))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, rc)
	return nil
}

func GetDristName(text string) string {
	return text[7:len(text)]
}

func GetNewDristName(text string) string {
	return text[7:len(text)]
}
