package pp

import (
	"fmt"
	. "github.com/k0kubun/palette"
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	object interface{}
	expect string
}

type Foo struct {
	Bar       int
	Hoge      string
	Hello     map[string]string
	HogeHoges []HogeHoge
}

type HogeHoge struct {
	Hell  string
	World int
}

var (
	testCases = []testCase{
		{true, BoldCyan("true")},
		{false, BoldCyan("false")},
		{int(4), BoldBlue("4")},
		{int8(8), BoldBlue("8")},
		{int16(16), BoldBlue("16")},
		{int32(32), BoldBlue("32")},
		{int64(64), BoldBlue("64")},
		{uint(4), BoldBlue("0x4")},
		{uint8(8), BoldBlue("0x8")},
		{uint16(16), BoldBlue("0x10")},
		{uint32(32), BoldBlue("0x20")},
		{uint64(64), BoldBlue("0x40")},
		{uintptr(128), BoldBlue("0x80")},
		{float32(2.23), BoldMagenta("2.23")},
		{float64(3.14), BoldMagenta("3.14")},
		{complex64(complex(3, -4)), BoldBlue("(3-4i)")},
		{complex128(complex(5, 6)), BoldBlue("(5+6i)")},
		{"string", BoldRed(`"`) + Red("string") + BoldRed(`"`)},
	}

	arr [3]int

	checkCases = []interface{}{
		map[string]int{"hell": 23, "world": 34},
		map[string]map[string]string{"s1": map[string]string{"v1": "m1", "va1": "me1"}, "si2": map[string]string{"v2": "m2"}},
		Foo{Bar: 1, Hoge: "a", Hello: map[string]string{"hel": "world", "a": "b"}, HogeHoges: []HogeHoge{HogeHoge{Hell: "a", World: 1}, HogeHoge{Hell: "bbb", World: 100}}},
		arr,
		[]string{"aaa", "bbb", "ccc"},
	}
)

func TestFormat(t *testing.T) {
	for _, test := range testCases {
		actual := fmt.Sprintf("%s", format(test.object))
		if test.expect != actual {
			v := reflect.ValueOf(test.object)
			t.Errorf("\nTestCase: %#v\nType: %s\nExpect: %# v\nActual: %# v\n", test.object, v.Kind(), test.expect, actual)
			continue
		}
		logResult(t, test.object, actual)
	}

	for _, object := range checkCases {
		actual := fmt.Sprintf("%s", format(object))
		logResult(t, object, actual)
	}
}

func logResult(t *testing.T, object interface{}, actual string) {
	if isMultiLine(actual) {
		t.Logf("%#v =>\n%s\n", object, actual)
	} else {
		t.Logf("%#v => %s\n", object, actual)
	}
}

func isMultiLine(text string) bool {
	return strings.Contains(text, "\n")
}
