package pp

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
	"unsafe"

	// Use fork until following PR is merged
	// https://github.com/mitchellh/colorstring/pull/3
	"github.com/k0kubun/colorstring"
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

type User struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LargeBuffer struct {
	Buf [1025]byte
}

type Circular struct {
	C *Circular
}

var c Circular = Circular{}

func init() {
	c.C = &c
}

var (
	testCases = []testCase{
		{nil, "[bold][cyan]nil"},
		{true, "[bold][cyan]true"},
		{false, "[bold][cyan]false"},
		{int(4), "[bold][blue]4"},
		{int8(8), "[bold][blue]8"},
		{int16(16), "[bold][blue]16"},
		{int32(32), "[bold][blue]32"},
		{int64(64), "[bold][blue]64"},
		{uint(4), "[bold][blue]0x4"},
		{uint8(8), "[bold][blue]0x8"},
		{uint16(16), "[bold][blue]0x10"},
		{uint32(32), "[bold][blue]0x20"},
		{uint64(64), "[bold][blue]0x40"},
		{uintptr(128), "[bold][blue]0x80"},
		{float32(2.23), "[bold][magenta]2.23"},
		{float64(3.14), "[bold][magenta]3.14"},
		{complex64(complex(3, -4)), "[bold][blue](3-4i)"},
		{complex128(complex(5, 6)), "[bold][blue](5+6i)"},
		{"string", `[bold][red]"[reset][red]string[reset][bold][red]"`},
		{[]string{}, "[][green]string[reset]{}"},
		{
			[]*Piyo{nil, nil}, `
			[]*pp.[green]Piyo[reset]{
			  (*pp.[green]Piyo[reset])([bold][cyan]nil[reset]),
			  (*pp.[green]Piyo[reset])([bold][cyan]nil[reset]),
			}
			`,
		},
		{
			&c, `
				&pp.[green]Circular[reset]{
				  [yellow]C[reset]: ...,
				}
			`,
		},
		{"日本\t語\x00", `[bold][red]"[reset][red]日本[reset][bold][magenta]\t[reset][red]語[reset][bold][magenta]\x00[reset][bold][red]"`},
		{
			time.Date(2015, time.February, 14, 22, 15, 0, 0, time.UTC),
			"[bold][blue]2015[reset]-[bold][blue]02[reset]-[bold][blue]14[reset] [bold][blue]22[reset]:[bold][blue]15[reset]:[bold][blue]00[reset] [bold][blue]UTC[reset]",
		},
		{
			LargeBuffer{}, `
			pp.[green]LargeBuffer[reset]{
			  [yellow]Buf[reset]: [[blue]1025[reset]][green]uint8[reset]{...},
			}
			`,
		},
	}

	arr [3]int
	tm  = time.Date(2015, time.January, 2, 0, 0, 0, 0, time.UTC)

	checkCases = []interface{}{
		map[string]int{"hell": 23, "world": 34},
		map[string]map[string]string{"s1": map[string]string{"v1": "m1", "va1": "me1"}, "si2": map[string]string{"v2": "m2"}},
		Foo{Bar: 1, Hoge: "a", Hello: map[string]string{"hel": "world", "a": "b"}, HogeHoges: []HogeHoge{HogeHoge{Hell: "a", World: 1}, HogeHoge{Hell: "bbb", World: 100}}},
		arr,
		[]string{"aaa", "bbb", "ccc"},
		make(chan bool, 10),
		func(a string, b float32) int { return 0 },
		&HogeHoge{},
		&Piyo{Field1: map[string]string{"a": "b", "cc": "dd"}, F2: &Foo{}, Fie3: 128},
		[]interface{}{1, 3},
		interface{}(1),
		HogeHoge{A: "test"},
		FooPri{Public: "hello", private: "world"},
		new(regexp.Regexp),
		unsafe.Pointer(new(regexp.Regexp)),
		"日本\t語\n\000\U00101234a",
		&tm,
		&User{Name: "k0kubun", CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()},
	}
)

func TestFormat(t *testing.T) {
	for _, test := range testCases {
		actual := fmt.Sprintf("%s", format(test.object))

		trimmed := strings.Replace(test.expect, "\t", "", -1)
		trimmed = strings.TrimPrefix(trimmed, "\n")
		trimmed = strings.TrimSuffix(trimmed, "\n")
		expect := colorstring.Color(trimmed)
		if expect != actual {
			v := reflect.ValueOf(test.object)
			t.Errorf("\nTestCase: %#v\nType: %s\nExpect: %# v\nActual: %# v\n", test.object, v.Kind(), expect, actual)
			return
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
