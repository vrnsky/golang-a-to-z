package api

const (
	// NewGameRoute is the path to create a new game.
	NewGameRoute = "/games"
)

// GameResponse contains the information about a game.
type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft byte    `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	WordLength   byte    `json:"word_length`
	Solution     string  `json:"solution,omitempty"`
	Status       string  `json:"status"`
}

// Guess is pair of word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
