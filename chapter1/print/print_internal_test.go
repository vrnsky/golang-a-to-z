package print

import "testing"

type testCase struct {
	lang language
	want string
}

var tests = map[string]testCase {
	"en": {
		lang: "en",
		want: "Hello World!",
	},
	"fr": {
		lang: "fr",
		want: "Bonjour le monde",
	},
	"ru": {
		lang: "ru",
		want: "Привет!",
	},
}

func ExamplePrint() {
	print()
	//Output:
	//Hello, World!
}

func TestGreet_Languages(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("Expected: %q, got %q", tc.want, got)
			}
		})
	}
}