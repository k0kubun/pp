package pp

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/mattn/go-colorable"
)

var out, defaultOut io.Writer
var outLock sync.Mutex

func init() {
	defaultOut = colorable.NewColorableStdout()
	out = defaultOut
}

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(out, formatAll(a)...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(out, format, formatAll(a)...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(out, formatAll(a)...)
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
	return fmt.Fprintln(w, formatAll(a)...)
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}

func Fatal(a ...interface{}) {
	fmt.Fprint(out, formatAll(a)...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, formatAll(a)...)
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	fmt.Fprintln(out, formatAll(a)...)
	os.Exit(1)
}

/*

Change Print* functions output to o
For example, you can limit output by ENV followings

	func init() {
		if os.Getenv("DEBUG") == "" {
			pp.SetDefaultOutput(ioutil.Discard)
		}
	}

*/
func SetDefaultOutput(o io.Writer) {
	outLock.Lock()
	out = o
	outLock.Unlock()
}

func GetDefaultOutput() io.Writer {
	return out
}

func ResetDefaultOutput() {
	outLock.Lock()
	out = defaultOut
	outLock.Unlock()
}

func formatAll(objects []interface{}) []interface{} {
	results := []interface{}{}
	for _, object := range objects {
		results = append(results, format(object))
	}
	return results
}
