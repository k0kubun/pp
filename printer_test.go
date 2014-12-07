package pp

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"unsafe"
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

type FooPri struct {
	Public  string
	private string
}

type Piyo struct {
	Field1 map[string]string
	F2     *Foo
	Fie3   int
}

type HogeHoge struct {
	Hell  string
	World int
	A     interface{}
}

var (
	testCases = []testCase{
		{nil, boldCyan("nil")},
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
		{"string", boldRed(`"`) + red("string") + boldRed(`"`)},
		{[]string{}, "[]" + green("string") + "{}"},
	}

	arr [3]int

	checkCases = []interface{}{
		map[string]int{"hell": 23, "world": 34},
		map[string]map[string]string{"s1": map[string]string{"v1": "m1", "va1": "me1"}, "si2": map[string]string{"v2": "m2"}},
		Foo{Bar: 1, Hoge: "a", Hello: map[string]string{"hel": "world", "a": "b"}, HogeHoges: []HogeHoge{HogeHoge{Hell: "a", World: 1}, HogeHoge{Hell: "bbb", World: 100}}},
		arr,
		[]string{"aaa", "bbb", "ccc"},
		make(chan bool, 10),
		unsafe.Pointer(uintptr(1)),
		func(a string, b float32) int { return 0 },
		&HogeHoge{},
		&Piyo{Field1: map[string]string{"a": "b", "cc": "dd"}, F2: &Foo{}, Fie3: 128},
		[]interface{}{1, 3},
		interface{}(1),
		HogeHoge{A: "test"},
		FooPri{Public: "hello", private: "world"},
		new(regexp.Regexp),
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
