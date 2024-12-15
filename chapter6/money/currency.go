package money

// ErrInvalidCurrencyCode is returnted when the currency to parse is not a standard 3-letter code.
const ErrInvalidCurrencyCode = Error("invalid currency code")

// ParserCurrency return the currency associated to a name and may return ErrInvalidCurrencyCode.
func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "OMR", "NTD":
		return Currency{code: code, precision: 2}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}

// Currency defines the code of a money and its decimal precision
type Currency struct {
	code      string
	precision byte
}

func NewCurrency(code string, precision byte) Currency {
	return Currency{
		code:      code,
		precision: precision,
	}
}

// String implements Stringer.
func (c Currency) String() string {
	return c.code
}
