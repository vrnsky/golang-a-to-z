package money

import (
	"fmt"
	"strconv"
	"strings"
)

// Decimal is capable of storing floating-point value.
type Decimal struct {
	// subunits is the amount of subunits. Multiply it by the precision to get the real value.
	subunits int64

	// precision - number of "subunits" in a unit, expressed as power of 10.
	precision byte
}

// ParseDecimal converts a string into its Decimal representation.
// It assumes there is up to one decimal separator, and that separator is '.' (full stop character).
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	// maxDecimal is the number of digits in a thousand billion
	const maxDecimal = 12

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))
	return Decimal{
		subunits:  subunits,
		precision: precision,
	}, nil
}
