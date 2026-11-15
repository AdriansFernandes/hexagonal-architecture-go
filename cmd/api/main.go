package main

import (
	"log"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/bootstrap"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/platform/httpserver"
)

func main() {
	handler := bootstrap.Bootstrap()
	server := httpserver.NewServer(":8080", handler)
	log.Fatal(server.ListenAndServe())
}
