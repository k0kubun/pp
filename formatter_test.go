package pp

import (
	"fmt"
	"testing"
)

type testCase struct {
	object interface{}
	expect string
}

var testCases = []testCase{
	{true, boldCyan("true")},
	{false, boldCyan("false")},
	{4, boldBlue("4")},
	{int8(8), boldBlue("8")},
	{int16(16), boldBlue("16")},
	{int32(32), boldBlue("32")},
	{int64(64), boldBlue("64")},
}

func TestFormat(t *testing.T) {
	for _, test := range testCases {
		actual := fmt.Sprintf("%s", format(test.object))
		t.Logf("%# v => %s\n", test.object, actual)
		if test.expect != actual {
			t.Errorf("\nTestCase: %# v\nExpect: %# v\nActual: %s\n", test.object, test.expect, actual)
		}
	}
}
