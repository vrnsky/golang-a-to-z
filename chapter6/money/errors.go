package money

const (
	// ErrInvalidDecimal is returned if the decimal is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")
	// ErrTooLarge is returned if the quantity is too large - this would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")
	ErrNegative = Error("Cannot be negative count")
)

// Error defines an error.
type Error string

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
