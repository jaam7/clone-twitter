package main

import (
	"log"

	"github.com/jalamar/clone-twitter/bd"
	"github.com/jalamar/clone-twitter/handlers"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("impossible connect to DB")
		return
	}

	handlers.Handlers()
}
