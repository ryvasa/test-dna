package main

import (
	"log"

	"github.com/ryvasa/test-dna/internal/app"
)

func main() {
	service := app.NewService()
	handler := app.NewHandler(service)
	router := app.NewRouter(handler)

	log.Fatal(router.Run(":8080"))
}
