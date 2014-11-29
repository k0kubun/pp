package pp

import (
	"fmt"
	"reflect"
)

var (
	colorByType = map[string]string{
		"bool":       "Cyan",
		"int":        "Blue",
		"int8":       "Blue",
		"int16":      "Blue",
		"int32":      "Blue",
		"int64":      "Blue",
		"uint":       "Blue",
		"uint8":      "Blue",
		"uint16":     "Blue",
		"uint32":     "Blue",
		"uint64":     "Blue",
		"uintptr":    "Blue",
		"float32":    "Magenta",
		"float64":    "Magenta",
		"complex64":  "Blue",
		"complex128": "Blue",
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

	if color, ok := colorByType[v.Kind().String()]; ok {
		fmt.Fprint(s, colorize(f.raw(), color))
		return
	}

	switch v.Kind() {
	case reflect.String:
		fmt.Fprint(s, boldRed("\"")+red(v.String())+boldRed("\""))
	default:
		fmt.Fprint(s, f.raw())
	}
}

func (f *formatter) raw() string{
	return fmt.Sprintf("%#v", f.object)
}
