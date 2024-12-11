package money

import (
	"errors"
	"testing"
)

func TestParseCurrency_Success(t *testing.T) {
	tt := map[string]struct {
		in       string
		expected Currency
		err      error
	}{
		"IRR in": {
			in:       "IRR",
			expected: Currency{code: "IRR", precision: 0},
			err:      nil,
		},
		"CNY in": {
			in:       "CNY",
			expected: Currency{code: "CNY", precision: 1},
			err:      nil,
		},
		"VND in": {
			in:       "VND",
			expected: Currency{code: "VND", precision: 1},
		},
		"MYR in": {
			in:       "MYR",
			expected: Currency{code: "MYR", precision: 2},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.in)
			if err != nil {
				t.Errorf("expected no error, got %s", err.Error())
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestParseCurrency_UnknowCurrency(t *testing.T) {
	_, err := ParseCurrency("INVALID")
	if !errors.Is(err, ErrInvalidCurrencyCode) {
		t.Errorf("expected %s, got %v", ErrInvalidCurrencyCode, err)
	}
}
