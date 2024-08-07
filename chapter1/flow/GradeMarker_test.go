package flow

import "testing"

func TestGetGrade(t *testing.T) {
	grade := GetGrade(80.00)
	if grade != "A" {
		t.Errorf("Got wrong degree for %f points. Expected %v, got: %v", 80.00, "A", grade)
	}

	grade = GetGrade(19.00)
	if grade != "F" {
		t.Errorf("Got wrong degree for %f points. Expected %v, got: %v", 20.00, "F", grade)
	}
}
