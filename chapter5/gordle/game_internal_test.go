package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in English": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), 0)
			got := g.ask()

			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("TOO LONG"),
			expected: errInvalidLengthMessage,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil, "no solution", 0)
			err := g.validateGuess(tc.word)

			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, tc.word)
			}
		})
	}
}
