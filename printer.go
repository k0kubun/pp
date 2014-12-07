package pp

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/tabwriter"
)

const (
	indentWidth = 2
)

func format(object interface{}) string {
	return newPrinter(object).String()
}

func newPrinter(object interface{}) *printer {
	buffer := bytes.NewBufferString("")
	tw := new(tabwriter.Writer)
	tw.Init(buffer, indentWidth, 0, 1, ' ', 0)

	return &printer{
		Buffer: buffer,
		tw:     tw,
		depth:  0,
		value:  reflect.ValueOf(object),
	}
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
	case reflect.Struct:
		p.printStruct()
	case reflect.Array, reflect.Slice:
		p.printSlice()
	case reflect.Chan:
		p.printf("(%s)(%s)", p.typeString(), p.pointerAddr())
	case reflect.Interface:
		p.printInterface()
	case reflect.Ptr:
		p.printPtr()
	case reflect.Func:
		p.printf("%s {...}", p.typeString())
	case reflect.UnsafePointer:
		p.printf("%s(%s)", p.typeString(), p.pointerAddr())
	case reflect.Invalid:
		p.print(p.nil())
	default:
		p.print(p.raw())
	}

	p.tw.Flush()
	return p.Buffer.String()
}

func (p *printer) print(text string) {
	fmt.Fprint(p.tw, text)
}

func (p *printer) printf(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	p.print(text)
}

func (p *printer) println(text string) {
	p.print(text + "\n")
}

func (p *printer) indentPrint(text string) {
	p.print(p.indent() + text)
}

func (p *printer) indentPrintf(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	p.indentPrint(text)
}

func (p *printer) colorPrint(text, color string) {
	p.print(colorize(text, color))
}

func (p *printer) printString() {
	p.colorPrint(`"`, "Red")
	p.colorPrint(p.value.String(), "red")
	p.colorPrint(`"`, "Red")
}

func (p *printer) printMap() {
	if p.value.Len() == 0 {
		p.printf("%s{}", p.typeString())
		return
	}

	p.println("{")
	p.indented(func() {
		keys := p.value.MapKeys()
		for i := 0; i < p.value.Len(); i++ {
			value := p.value.MapIndex(keys[i])
			p.indentPrintf("%s:\t%s,\n", p.format(keys[i]), p.format(value))
		}
	})
	p.indentPrint("}")
}

func (p *printer) printStruct() {
	p.println(p.typeString() + "{")
	p.indented(func() {
		for i := 0; i < p.value.NumField(); i++ {
			field := yellow(p.value.Type().Field(i).Name)
			value := p.value.Field(i)
			p.indentPrintf("%s:\t%s,\n", field, p.format(value))
		}
	})
	p.indentPrint("}")
}

func (p *printer) printSlice() {
	if p.value.Len() == 0 {
		p.printf("%s{}", p.typeString())
		return
	}

	p.println(p.typeString() + "{")
	p.indented(func() {
		for i := 0; i < p.value.Len(); i++ {
			p.indentPrintf("%s,\n", p.format(p.value.Index(i)))
		}
	})
	p.indentPrint("}")
}

func (p *printer) printInterface() {
	e := p.value.Elem()
	if e.Kind() == reflect.Invalid {
		p.print(p.nil())
	} else if e.IsValid() {
		p.print(p.format(e))
	} else {
		p.printf("%s(%s)", p.typeString(), p.nil())
	}
}

func (p *printer) printPtr() {
	if p.value.Elem().IsValid() {
		p.printf("&%s", p.format(p.value.Elem()))
	} else {
		p.printf("(%s)(%s)", p.typeString(), p.nil())
	}
}

func (p *printer) pointerAddr() string {
	return boldBlue(fmt.Sprintf("%#v", p.value.Pointer()))
}

func (p *printer) typeString() string {
	prefix := ""
	t := p.value.Type().String()

	if p.matchRegexp(t, `^\[\].+$`) {
		prefix = "[]"
		t = t[2:]
	}

	if p.matchRegexp(t, `^\[\d\].+$`) {
		num := regexp.MustCompile(`\d`).FindString(t)
		prefix = fmt.Sprintf("[%s]", blue(num))
		t = t[2+len(num):]
	}

	if p.matchRegexp(t, `^[^\.]+\.[^\.]+$`) {
		ts := strings.Split(t, ".")
		t = fmt.Sprintf("%s.%s", ts[0], green(ts[1]))
	} else {
		t = green(t)
	}
	return prefix + t
}

func (p *printer) matchRegexp(text, exp string) bool {
	return regexp.MustCompile(exp).MatchString(text)
}

func (p *printer) indented(proc func()) {
	p.depth++
	proc()
	p.depth--
}

func (p *printer) raw() string {
	// Some value causes panic when Interface() is called.
	switch p.value.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%#v", p.value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%#v", p.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%#v", p.value.Uint())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%#v", p.value.Complex())
	default:
		return fmt.Sprintf("%#v", p.value.Interface())
	}
}

func (p *printer) nil() string {
	return boldCyan("nil")
}

func (p *printer) format(object interface{}) string {
	pp := newPrinter(object)
	pp.depth = p.depth
	if value, ok := object.(reflect.Value); ok {
		pp.value = value
	}
	return pp.String()
}

func (p *printer) indent() string {
	return strings.Repeat("\t", p.depth)
}
