package main

import (
	"log"

	"github.com/jainparan/backend-server/src/app"
)

func main() {
	err := app.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
