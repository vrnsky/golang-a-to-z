package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const wordLength = 5

type Game struct {
	reader *bufio.Reader
}

func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}
	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	guess := g.ask()
	fmt.Printf("Your guess is: %s\n", string(guess))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", wordLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
		}

		guess := []rune(string(playerInput))

		if len(guess) != wordLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with with Gordle's solution! Expected %d character, got %d.", wordLength, len(guess))
		} else {
			return guess
		}
	}
}
