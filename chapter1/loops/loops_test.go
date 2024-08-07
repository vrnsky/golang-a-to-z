package loops

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 3)
	if result != 6 {
		t.Errorf("Expected sum from 1 to 3 is %d, but got %d", 6, result)
	}
}
