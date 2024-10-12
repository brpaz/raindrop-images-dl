package main

import (
	"log"

	"github.com/brpaz/raindrop-images-dl/internal/app"
)

func main() {
	app := app.New()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
