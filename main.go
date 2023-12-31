package main

import (
	"log"
    "fmt"
)

func main() {
	server := NewServer(":3000")

	// server.Get("/", server.AddMiddleware(HandleRoot, Logger))
    server.All("/", server.AddMiddleware(HandleRoot, Logger))
	server.Post("/api", server.AddMiddleware(HandleApi, Logger, CheckAuth))
    server.Post("/create", server.AddMiddleware(PostRequest, Logger))

    fmt.Printf("%+v\n", server.router.rules)
	log.Fatal(server.Listen())
}
