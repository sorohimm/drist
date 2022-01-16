package main

import (
	"io"
	"os"
)

func SaveDrist(rc io.ReadCloser, fn string) error {
  file, err := os.Create("./memes/" + fn)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = io.Copy(file, rc)
  return err
}

