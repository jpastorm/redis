package main

import (
	"github.com/jpastorm/redis/cmd/client/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}