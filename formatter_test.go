package pp

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	object interface{}
	expect string
}

var (
	testCases = []testCase{
		{true, boldCyan("true")},
		{false, boldCyan("false")},
		{int(4), boldBlue("4")},
		{int8(8), boldBlue("8")},
		{int16(16), boldBlue("16")},
		{int32(32), boldBlue("32")},
		{int64(64), boldBlue("64")},
		{uint(4), boldBlue("0x4")},
		{uint8(8), boldBlue("0x8")},
		{uint16(16), boldBlue("0x10")},
		{uint32(32), boldBlue("0x20")},
		{uint64(64), boldBlue("0x40")},
		{uintptr(128), boldBlue("0x80")},
		{float32(2.23), boldMagenta("2.23")},
		{float64(3.14), boldMagenta("3.14")},
		{complex64(complex(3, -4)), boldBlue("(3-4i)")},
		{complex128(complex(5, 6)), boldBlue("(5+6i)")},
		{"string", boldRed("\"") + red("string") + boldRed("\"")},
	}

	checkCases = []interface{}{
		map[string]int{"hello": 23, "world": 34},
		map[string]map[string]string{"s1": map[string]string{"v1": "m1"}, "s2": map[string]string{"v2": "m2"}},
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
