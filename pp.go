package pp

import (
	"fmt"
)

func Print(object interface{}) {
	fmt.Print(format(object))
}
