package pp

import (
	"bytes"
	"fmt"
	"reflect"
	"text/tabwriter"
)

const (
	indentWidth = 2
)

func format(object interface{}) string {
	buffer := bytes.NewBufferString("")
	tw := new(tabwriter.Writer)
	tw.Init(buffer, indentWidth, 0, 1, ' ', 0)

	p := &printer{
		Buffer: buffer,
		tw:     tw,
		depth:  0,
		value:  reflect.ValueOf(object),
	}
	return p.String()
}

type printer struct {
	*bytes.Buffer
	tw    *tabwriter.Writer
	depth int
	value reflect.Value
}

func (p *printer) String() string {
	switch p.value.Kind() {
	case reflect.Bool:
		p.colorPrint("Cyan")
	default:
		p.print(p.raw())
	}
	return p.Buffer.String()
}

func (p *printer) print(text string) {
	fmt.Fprint(p.Buffer, text)
}

func (p *printer) colorPrint(color string) {
	p.print(colorize(p.raw(), color))
}

func (p *printer) raw() string {
	return fmt.Sprintf("%#v", p.value.Interface())
}
