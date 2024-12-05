package money

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"2 decimal digits": {
			decimal: "1.52",
			expected: Decimal{
				subunits:  152,
				precision: 2,
			},
			err: nil,
		},
		"no decimal digits": {
			decimal:  "1",
			expected: Decimal{1, 0},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			decimal:  "1.02",
			expected: Decimal{102, 2},
			err:      nil,
		},
		"multiple of 10": {
			decimal:  "150",
			expected: Decimal{150, 0},
			err:      nil,
		},
		"invalid decimal format": {
			decimal: "65.packet",
			err:     ErrInvalidDecimal,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("execpted %v, got %v", tc.expected, got)
			}
		})
	}
}
