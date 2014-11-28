package pp

import (
	"fmt"
)

func format(object interface{}) *formatter {
	return &formatter{object}
}

type formatter struct {
	object interface{}
}

func (f *formatter) Format(s fmt.State, c rune) {
	fmt.Println(color("test", "Red"))
}
