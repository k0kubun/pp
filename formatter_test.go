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
