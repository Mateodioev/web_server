package main

import (
	"log"
)

func main() {
	server := NewServer(":3000")

	server.Get("/", server.AddMiddleware(HandleRoot, Logger))
	server.Post("/api", server.AddMiddleware(HandleApi, Logger, CheckAuth))

	log.Fatal(server.Listen())
}
