package main

import (
	"gopet/server"
	"gopet/service"
	"gopet/storage"
	"log"
)

func main() {
	store, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}

	svc := service.New(store)
	srv := server.New(svc)

	log.Println("Server started on :8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
