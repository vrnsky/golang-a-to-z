package session

// A GameId represents the ID of a game.
type GameId string

// Status is the current status of the game and tells what operations can be made on it.
type Status string

// A Guess is pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string
	Feedback string
}

// Game contains the information about a game.
type Game struct {
	ID           GameId
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}
