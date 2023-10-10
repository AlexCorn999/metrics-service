package main

import (
	"log"

	"github.com/AlexCorn999/metrics-service/internal/transport/apiserver"
)

func main() {
	server := apiserver.NewAPIServer()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
