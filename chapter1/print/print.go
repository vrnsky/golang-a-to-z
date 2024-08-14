package print

import "fmt"

func print() {
	fmt.Println("Hello, World!")
}

type language string

var phaseBook = map[language]string {
	"en": "Hello World!",
	"fr": "Bonjour le monde",
	"ru": "Привет!",
	"": "",

}
func greet(language language) string {
	result, ok := phaseBook[language]
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", language)
	}
	return result
}