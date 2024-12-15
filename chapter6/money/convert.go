package money

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	convertedValue := applyExchangeRate(amount, to, ExchangeRate{subunits: 2, precision: 0})

	if err := convertedValue.Validate(); err != nil {
		return Amount{}, err
	}
	return convertedValue, nil
}

type ExchangeRate Decimal

func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted := multiply(a.quantity, rate)

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}
}

// Validate checks if the amount is valid according to currency rules
func (a Amount) Validate() error {
	if a.quantity.subunits < 0 {
		return ErrNegative
	}
	if a.quantity.precision > a.currency.precision {
		return ErrTooPrecise
	}
	if a.quantity.subunits >= pow10(19) {
		return ErrTooLarge
	}
	return nil
}

func multiply(d Decimal, r ExchangeRate) Decimal {
	return Decimal{
		subunits:  d.subunits * r.subunits,
		precision: d.precision + r.precision,
	}
}
