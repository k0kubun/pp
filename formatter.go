package pp

type formatter struct {
	object interface{}
}

func (f *formatter) String() string {
	return "test\n"
}
