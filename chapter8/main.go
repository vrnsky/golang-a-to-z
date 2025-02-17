package main

import (
	"fmt"
	"golang-a-to-z/chapter8/handlers"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.Mux())

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
