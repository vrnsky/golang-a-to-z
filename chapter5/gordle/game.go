package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

var errInvalidLengthMessage = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

type Game struct {
	reader *bufio.Reader
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidLengthMessage)
	}
	return nil
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := []rune(string(playerInput))
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	guess := g.ask()
	fmt.Printf("Your guess is: %s\n", string(guess))
}

func New(playerInput io.Reader) *Game {
	game := &Game{
		reader: bufio.NewReader(playerInput),
	}
	return game
}
