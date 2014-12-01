package pp

import (
	"errors"
	"fmt"
	"io"
)

func Print(a ...interface{}) (n int, err error) {
	return fmt.Print(formatAll(a)...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, formatAll(a)...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(formatAll(a)...)
}

func Sprint(a ...interface{}) string {
	return fmt.Sprint(formatAll(a)...)
}

func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, formatAll(a)...)
}

func Sprintln(a ...interface{}) string {
	return fmt.Sprintln(formatAll(a)...)
}

func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, formatAll(a)...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, formatAll(a)...)
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return Fprintln(w, formatAll(a)...)
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}

func formatAll(objects []interface{}) []interface{} {
	results := []interface{}{}
	for _, object := range objects {
		results = append(results, format(object))
	}
	return results
}
