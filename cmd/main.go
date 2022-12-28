package main

import (
	"log"

	"github.com/suvrick/kiss/internal/app"
)

func main() {
	app := app.App{}
	if err := app.Run(); err != nil {
		log.Println(err.Error())
	}
}
