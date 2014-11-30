package pp

import (
	"bytes"
	"fmt"
	"github.com/k0kubun/palette"
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
		p.colorPrint(p.raw(), "Cyan")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		p.colorPrint(p.raw(), "Blue")
	case reflect.Float32, reflect.Float64:
		p.colorPrint(p.raw(), "Magenta")
	case reflect.String:
		p.printString()
	case reflect.Map:
		p.printMap()
	default:
		p.print(p.raw())
	}
	return p.Buffer.String()
}

func (p *printer) print(text string) {
	fmt.Fprint(p.Buffer, text)
}

func (p *printer) println(text string) {
	p.print(text+"\n")
}

func (p *printer) colorPrint(text, color string) {
	p.print(palette.Colorize(text, color))
}

func (p *printer) raw() string {
	return fmt.Sprintf("%#v", p.value.Interface())
}

func (p *printer) printString() {
	p.colorPrint(`"`, "Red")
	p.colorPrint(p.value.String(), "red")
	p.colorPrint(`"`, "Red")
}

func (p *printer) printMap() {
	p.println("{")
	keys := p.value.MapKeys()

	for i := 0; i < p.value.Len(); i++ {
		key := keys[i].Interface()
		value := p.value.MapIndex(keys[i]).Interface()
		fmt.Fprintf(p.tw, "\t%s:\t%s,\n", format(key), format(value))
	}
	p.tw.Flush()
	p.println("}")
}
