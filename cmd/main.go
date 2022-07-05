package main

import (
	"log"

	server "ascii-art-web/internal/app"
)

func main() {
	if err := server.App(); err != nil {
		log.Println(err.Error())
		return
	}
}
