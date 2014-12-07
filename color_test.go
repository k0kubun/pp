package pp

import (
	"testing"
)

type colorTest struct {
	input  string
	result string
}

var (
	expects = []colorTest{
		{"black", "\x1b[30mpalette\x1b[0m"},
		{"red", "\x1b[31mpalette\x1b[0m"},
		{"green", "\x1b[32mpalette\x1b[0m"},
		{"yellow", "\x1b[33mpalette\x1b[0m"},
		{"blue", "\x1b[34mpalette\x1b[0m"},
		{"magenta", "\x1b[35mpalette\x1b[0m"},
		{"cyan", "\x1b[36mpalette\x1b[0m"},
		{"white", "\x1b[37mpalette\x1b[0m"},
		{"Black", "\x1b[30m\x1b[1mpalette\x1b[0m"},
		{"Red", "\x1b[31m\x1b[1mpalette\x1b[0m"},
		{"Green", "\x1b[32m\x1b[1mpalette\x1b[0m"},
		{"Yellow", "\x1b[33m\x1b[1mpalette\x1b[0m"},
		{"Blue", "\x1b[34m\x1b[1mpalette\x1b[0m"},
		{"Magenta", "\x1b[35m\x1b[1mpalette\x1b[0m"},
		{"Cyan", "\x1b[36m\x1b[1mpalette\x1b[0m"},
		{"White", "\x1b[37m\x1b[1mpalette\x1b[0m"},
	}
)

func TestColorize(t *testing.T) {
	for _, test := range expects {
		expect(t, test.input, test.result)
	}

	t.Logf(black("black"))
	t.Logf(red("red"))
	t.Logf(green("green"))
	t.Logf(yellow("yellow"))
	t.Logf(blue("blue"))
	t.Logf(magenta("magenta"))
	t.Logf(cyan("cyan"))
	t.Logf(white("white"))
	t.Logf(boldBlack("Black"))
	t.Logf(boldRed("Red"))
	t.Logf(boldGreen("Green"))
	t.Logf(boldYellow("Yellow"))
	t.Logf(boldBlue("Blue"))
	t.Logf(boldMagenta("Magenta"))
	t.Logf(boldCyan("Cyan"))
	t.Logf(boldWhite("White"))
}

func expect(t *testing.T, input, result string) {
	actual := colorize("palette", input)
	if actual != result {
		t.Errorf("Expected: %#v, Actual: %#v", result, actual)
	} else {
		t.Logf("%s => %s", input, actual)
	}
}
