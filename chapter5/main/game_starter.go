package main

import (
	"golang-a-to-z/chapter5/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
