package main

import (
	"os"

	"github.com/segmentationfaulter/todo-cli-app/cmd"
	"github.com/segmentationfaulter/todo-cli-app/storage"
)

func main() {
	path, err := storage.GetDataFilePath()
	if err != nil {
		panic(err)
	}

	_, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
