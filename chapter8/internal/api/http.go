package api

const (
	// GameId is the name of the field that stores the game's identifier.
	GameId = "id"

	// NewGameRoute is the path to create a new game.
	NewGameRoute = "/games"

	// GetStatusRoute is the path to get the status of game identified by its id.
	GetStatusRoute = "/games/{" + GameId + "}"

	GuessRoute = "/games/{" + GameId + "}"
)

// GameResponse contains the information about a game.
type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft byte    `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	WordLength   byte    `json:"word_length"`
	Solution     string  `json:"solution,omitempty"`
	Status       string  `json:"status"`
}

// Guess is pair of word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}

// GuessRequest is the structure of the message used when submitting a guess.
type GuessRequest struct {
	Guess string `json:"guess"`
}
