package pp

import (
	"testing"
)

type colorTest struct {
	input  string
	color  FlagSet
	result string
}

var tests = []colorTest{
	colorTest{
		"Hello",
		Blue | BackRed,
		"\x1b[34m\x1b[41mHello\x1b[0m",
	},
	colorTest{
		"This is me",
		Magenta | BackWhite,
		"\x1b[35m\x1b[47mThis is me\x1b[0m",
	},
	colorTest{
		"How are you",
		Cyan,
		"\x1b[36mHow are you\x1b[0m",
	},
	colorTest{
		"This isssss gettting boring",
		BackRed,
		"\x1b[41mThis isssss gettting boring\x1b[0m",
	},
	colorTest{
		"DONE",
		Bold,
		"\x1b[1mDONE\x1b[0m",
	},
}

func TestColorize(t *testing.T) {
	for _, test := range tests {
		if output := colorize(test.input, test.color); output != test.result {
			t.Errorf("Expected %q, got %q", test.result, output)
		}
	}
}
