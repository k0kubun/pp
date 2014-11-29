package pp

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	object interface{}
	expect string
}

var testCases = []testCase{
	{true, boldCyan("true")},
	{false, boldCyan("false")},
	{1, boldBlue("1")},
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
	{1.41, boldMagenta("1.41")},
	{float32(2.23), boldMagenta("2.23")},
	{float64(3.14), boldMagenta("3.14")},
}

func TestFormat(t *testing.T) {
	for _, test := range testCases {
		actual := fmt.Sprintf("%s", format(test.object))
		if test.expect != actual {
			v := reflect.ValueOf(test.object)
			t.Errorf("\nTestCase: %# v\nType: %s\nExpect: %# v\nActual: %s\n", test.object, v.Kind(), test.expect, actual)
		} else {
			t.Logf("%#v => %s\n", test.object, actual)
		}
	}
}
