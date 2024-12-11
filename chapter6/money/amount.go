package money

const (
	// ErrToPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("Qunaitity is too precise")
)

type Amount struct {
	quantity Decimal
	currency Currency
}

// NewAmount return an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	quantity.precision = currency.precision
	return Amount{quantity: quantity, currency: currency}, nil
}
