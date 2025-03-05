package session

// Error is sentinel error for the domain.
type Error string

func (d Error) Error() string {
	return string(d)
}

const (
	ErrGameOver = Error("😞 game over")
)
