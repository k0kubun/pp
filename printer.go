package pp

import (
	"bytes"
	"io"
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
		Writer: buffer,
		tw:     tw,
		depth:  0,
	}
	return p.String()
}

type printer struct {
	io.Writer
	tw    *tabwriter.Writer
	depth int
}

func (p *printer) String() string {
	return ""
}
