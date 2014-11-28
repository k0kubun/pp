package pp

func format(object interface{}) *formatter {
	return &formatter{object}
}

type formatter struct {
	object interface{}
}

func (f *formatter) String() string {
	return color("test\n", "red")
}
