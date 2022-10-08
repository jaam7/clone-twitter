package main

import "log"

import(
	"github.com/jaam7/clone-twitter/bd"
)

func main() {
	if bd.checkConnection() == false {
		log.Fatal("impossible connect to DB")
	}
}
