package pp

import (
	"fmt"
	"reflect"
)

var (
	funcByType = map[string]func(interface{}) string{
		"bool":       colorFormatter("Cyan"),
		"int":        colorFormatter("Blue"),
		"int8":       colorFormatter("Blue"),
		"int16":      colorFormatter("Blue"),
		"int32":      colorFormatter("Blue"),
		"int64":      colorFormatter("Blue"),
		"uint":       colorFormatter("Blue"),
		"uint8":      colorFormatter("Blue"),
		"uint16":     colorFormatter("Blue"),
		"uint32":     colorFormatter("Blue"),
		"uint64":     colorFormatter("Blue"),
		"uintptr":    colorFormatter("Blue"),
		"float32":    colorFormatter("Magenta"),
		"float64":    colorFormatter("Magenta"),
		"complex64":  colorFormatter("Blue"),
		"complex128": colorFormatter("Blue"),
		"string":     formatString,
	}
)

func format(object interface{}) *formatter {
	return &formatter{object}
}

type formatter struct {
	object interface{}
}

func (f *formatter) String() string {
	return fmt.Sprint(f.object)
}

func (f *formatter) Format(s fmt.State, c rune) {
	v := reflect.ValueOf(f.object)

	if fc, ok := funcByType[v.Kind().String()]; ok {
		fmt.Fprint(s, fc(f.object))
	} else {
		fmt.Fprint(s, fmt.Sprintf("%#v", f.object))
	}
}

func colorFormatter(color string) func(interface{}) string {
	return func(object interface{}) string {
		raw := fmt.Sprintf("%#v", object)
		return colorize(raw, color)
	}
}

func formatString(object interface{}) string {
	return boldRed("\"") + red(fmt.Sprintf("%s", object)) + boldRed("\"")
}
