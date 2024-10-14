package main

import (
	"log"

	delivery "github.com/andreanpradanaa/api-gateway-repo/internal/delivery/rest"
)

func main() {
	httpServer := delivery.NewHttpHandler()
	err := httpServer.Start()
	if err != nil {
		log.Fatalf("cannot create run HTTP server: %s", err)
	}

	log.Print("listen at port:8080")
}
