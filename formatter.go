package pp

import (
	"fmt"
	"reflect"
)

var (
	funcByType = map[string]func(reflect.Value) string{
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
		"map":        formatMap,
	}
)

func format(object interface{}) *formatter {
	return &formatter{object}
}

type formatter struct {
	object interface{}
}

func (f *formatter) String() string {
	return fmt.Sprint(f)
}

func (f *formatter) Format(s fmt.State, c rune) {
	v := reflect.ValueOf(f.object)

	if fc, ok := funcByType[v.Kind().String()]; ok {
		fmt.Fprint(s, fc(v))
	} else {
		fmt.Fprint(s, fmt.Sprintf("%#v", f.object))
	}
}

func colorFormatter(color string) func(reflect.Value) string {
	return func(v reflect.Value) string {
		raw := fmt.Sprintf("%#v", v.Interface())
		return colorize(raw, color)
	}
}

func formatString(v reflect.Value) string {
	return boldRed(`"`) + red(v.String()) + boldRed(`"`)
}

func formatMap(v reflect.Value) string {
	result := "{\n"
	keys := v.MapKeys()
	for i := 0; i < v.Len(); i++ {
		key := keys[i]
		result += "  "
		result += format(key.Interface()).String()
		result += ": "
		result += format(v.MapIndex(key).Interface()).String()
		result += ",\n"
	}
	result += "}"
	return result
}
