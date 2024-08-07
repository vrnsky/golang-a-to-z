package arithemtics

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 10)
	if result != 20 {
		t.Errorf("Result of add function is incorrect, got %d, expected %d", result, 20)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(10, 2)
	if result != 20 {
		t.Errorf("Result of multiply function is incorrect, got %d, expected %d", result, 20)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 2)
	if result != 8 {
		t.Errorf("Result of multiply function is incorrect, got %d, expected %d", result, 20)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(20, 2)
	if result != 10 {
		t.Errorf("Result of multiply function is incorrect, got %d, expected %d", result, 20)
	}
}

func TestDivideByZeroShouldPanic(t *testing.T) {
	defer func() { _ = recover() }()
	Divide(100, 0)

	t.Errorf("Divison by zero did not panic!")
}
