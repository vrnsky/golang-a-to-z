package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

const solutionLength = 5

var errInvalidLengthMessage = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidLengthMessage)
	}
	return nil
}

// splitToUppercaseCharacter is naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacter(input string) []rune {
	return []rune(strings.ToUpper(input))
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
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🔥You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
		fmt.Printf("🗿You've lost! The solution was: %s.\n", string(g.solution))
	}
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	game := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacter(solution),
		maxAttempts: maxAttempts,
	}
	return game
}
