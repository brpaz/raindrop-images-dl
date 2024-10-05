package main

import (
	"log"

	"github.com/brpaz/raindrop-images-dl/cmd"
)

func main() {
	app := cmd.NewApp()
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
