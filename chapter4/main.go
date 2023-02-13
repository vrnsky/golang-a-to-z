package main

import (
	"bufio"
	"os"
	"vrnsky.io/m/v2/gordle"
)

func main() {
	g := gordle.New(bufio.NewReader(os.Stdin))
	g.Play()
}
