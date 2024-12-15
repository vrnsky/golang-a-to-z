package money

const (
	// ErrToPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("Qunaitity is too precise")
	ErrTooLarger  = Error("Quantity is too large")
)

type Amount struct {
	quantity Decimal
	currency Currency
}

func validate(a Amount) error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}
	return nil
}

// NewAmount return an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	quantity.precision = currency.precision
	return Amount{quantity: quantity, currency: currency}, nil
}

// String implements stringer.
func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.code
}
