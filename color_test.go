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

	t.Logf(Black("black"))
	t.Logf(Red("red"))
	t.Logf(Green("green"))
	t.Logf(Yellow("yellow"))
	t.Logf(Blue("blue"))
	t.Logf(Magenta("magenta"))
	t.Logf(Cyan("cyan"))
	t.Logf(White("white"))
	t.Logf(BoldBlack("Black"))
	t.Logf(BoldRed("Red"))
	t.Logf(BoldGreen("Green"))
	t.Logf(BoldYellow("Yellow"))
	t.Logf(BoldBlue("Blue"))
	t.Logf(BoldMagenta("Magenta"))
	t.Logf(BoldCyan("Cyan"))
	t.Logf(BoldWhite("White"))
}

func expect(t *testing.T, input, result string) {
	actual := Colorize("palette", input)
	if actual != result {
		t.Errorf("Expected: %#v, Actual: %#v", result, actual)
	} else {
		t.Logf("%s => %s", input, actual)
	}
}
