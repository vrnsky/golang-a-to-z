package print

import "testing"

func ExamplePrint() {
	print()
	//Output:
	//Hello, World!
}

func TestGreet(t *testing.T) {
	result := greet()
	if result != "Golang is great!" {
		t.Errorf("Expected %v, got: %v", "Golang is great!", result)
	}
}