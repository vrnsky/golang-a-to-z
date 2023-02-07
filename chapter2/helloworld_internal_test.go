package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	expectedGreeting := "Hello world"

	greeting := greet("en")

	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	expectedGreeting := "Bonjour le monde"

	greeting := greet("fr")

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	expectedGreeting := ""

	greeting := greet("akk")

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}

}
