package pp

import (
	"fmt"
	"reflect"
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

	switch v.Kind() {
	case reflect.Bool:
		fmt.Fprint(s, boldCyan(fmt.Sprintf("%#v", v.Bool())))
	default:
		fmt.Fprint(s, fmt.Sprintf("%#v", f.object))
	}
}
