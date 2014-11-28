package pp

import (
	"fmt"
)

func Print(object interface{}) {
	fmt.Print(format(object))
}

func format(object interface{}) *formatter {
	return &formatter{object}
}
