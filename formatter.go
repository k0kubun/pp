package pp

import (
	"fmt"
	"reflect"
)

var (
	colorByType = map[string]string{
		"bool":    "Cyan",
		"int":     "Blue",
		"int8":    "Blue",
		"int16":   "Blue",
		"int32":   "Blue",
		"int64":   "Blue",
		"uint":    "Blue",
		"uint8":   "Blue",
		"uint16":  "Blue",
		"uint32":  "Blue",
		"uint64":  "Blue",
		"uintptr": "Blue",
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
		raw := fmt.Sprintf("%#v", f.object)
		fmt.Fprint(s, colorize(raw, color))
		return
	}

	fmt.Fprint(s, fmt.Sprintf("%#v", f.object))
}
