package pp

import (
	"fmt"
	"testing"
)

type colorTest struct {
	input  string
	result string
}

func TestColorize(t *testing.T) {
	fmt.Println(colorize("Hello", Blue|BackBlue|Bold))

	color := colorTest{
		input:  "Hi",
		result: "Let's hope this works",
	}

	Print(color)
}
