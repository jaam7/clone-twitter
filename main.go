package main

import (
	"github.com/jalamar/clone-twitter/bd"

	"log"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("impossible connect to DB")
	}
}
