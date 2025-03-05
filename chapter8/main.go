package main

import (
	"golang-a-to-z/chapter8/internal/handlers"
	"golang-a-to-z/chapter8/internal/repository"
	"log"
	"net/http"
)

func main() {
	db := repository.New()

	addr := ":9090"

	log.Print("Listening on ", addr, "...")

	err := http.ListenAndServe(addr, handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
