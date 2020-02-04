package pp

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/mattn/go-colorable"
)

var (
	defaultOut           = colorable.NewColorableStdout()
	defaultWithLineInfo  = false
	defaultPrettyPrinter = New()
)

type PrettyPrinter struct {
	out           io.Writer
	currentScheme ColorScheme
	// WithLineInfo add file name and line information to output
	// call this function with care, because getting stack has performance penalty
	WithLineInfo    bool
	outLock         sync.Mutex
	maxDepth        int
	coloringEnabled bool
}

// New creates a new PrettyPrinter that can be used to pretty print values
func New() *PrettyPrinter {
	return &PrettyPrinter{
		out:             defaultOut,
		currentScheme:   defaultScheme,
		WithLineInfo:    defaultWithLineInfo,
		maxDepth:        -1,
		coloringEnabled: true,
	}
}

// Print prints given arguments.
func (pp *PrettyPrinter) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(pp.out, pp.formatAll(a)...)
}

// Printf prints a given format.
func (pp *PrettyPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(pp.out, format, pp.formatAll(a)...)
}

// Println prints given arguments with newline.
func (pp *PrettyPrinter) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(pp.out, pp.formatAll(a)...)
}

// Sprint formats given arguemnts and returns the result as string.
func (pp *PrettyPrinter) Sprint(a ...interface{}) string {
	return fmt.Sprint(pp.formatAll(a)...)
}

// Sprintf formats with pretty print and returns the result as string.
func (pp *PrettyPrinter) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, pp.formatAll(a)...)
}

// Sprintln formats given arguemnts with newline and returns the result as string.
func (pp *PrettyPrinter) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(pp.formatAll(a)...)
}

// Fprint prints given arguments to a given writer.
func (pp *PrettyPrinter) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, pp.formatAll(a)...)
}

// Fprintf prints format to a given writer.
func (pp *PrettyPrinter) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, pp.formatAll(a)...)
}

// Fprintln prints given arguments to a given writer with newline.
func (pp *PrettyPrinter) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, pp.formatAll(a)...)
}

// Errorf formats given arguments and returns it as error type.
func (pp *PrettyPrinter) Errorf(format string, a ...interface{}) error {
	return errors.New(pp.Sprintf(format, a...))
}

// Fatal prints given arguments and finishes execution with exit status 1.
func (pp *PrettyPrinter) Fatal(a ...interface{}) {
	fmt.Fprint(pp.out, pp.formatAll(a)...)
	os.Exit(1)
}

// Fatalf prints a given format and finishes execution with exit status 1.
func (pp *PrettyPrinter) Fatalf(format string, a ...interface{}) {
	fmt.Fprintf(pp.out, format, pp.formatAll(a)...)
	os.Exit(1)
}

// Fatalln prints given arguments with newline and finishes execution with exit status 1.
func (pp *PrettyPrinter) Fatalln(a ...interface{}) {
	fmt.Fprintln(pp.out, pp.formatAll(a)...)
	os.Exit(1)
}

func (pp *PrettyPrinter) SetColoringEnabled(enabled bool) {
	pp.coloringEnabled = enabled
}

// SetOutput sets pp's output
func (pp *PrettyPrinter) SetOutput(o io.Writer) {
	pp.outLock.Lock()
	pp.out = o
	pp.outLock.Unlock()
}

// GetOutput returns pp's output.
func (pp *PrettyPrinter) GetOutput() io.Writer {
	return pp.out
}

// ResetOutput sets pp's output back to the default output
func (pp *PrettyPrinter) ResetOutput() {
	pp.outLock.Lock()
	pp.out = defaultOut
	pp.outLock.Unlock()
}

// SetColorScheme takes a colorscheme used by all future Print calls.
func (pp *PrettyPrinter) SetColorScheme(scheme ColorScheme) {
	scheme.fixColors()
	pp.currentScheme = scheme
}

// ResetColorScheme resets colorscheme to default.
func (pp *PrettyPrinter) ResetColorScheme() {
	pp.currentScheme = defaultScheme
}

func (pp *PrettyPrinter) formatAll(objects []interface{}) []interface{} {
	results := []interface{}{}

	// fix for backwards capability
	withLineInfo := pp.WithLineInfo
	if pp == defaultPrettyPrinter {
		withLineInfo = WithLineInfo
	}

	if withLineInfo {
		_, fn, line, _ := runtime.Caller(2) // 2 because current Caller is pp itself
		results = append(results, fmt.Sprintf("%s:%d\n", fn, line))
	}

	for _, object := range objects {
		results = append(results, pp.format(object))
	}
	return results
}

// Print prints given arguments.
func Print(a ...interface{}) (n int, err error) {
	return defaultPrettyPrinter.Print(a...)
}

// Printf prints a given format.
func Printf(format string, a ...interface{}) (n int, err error) {
	return defaultPrettyPrinter.Printf(format, a...)
}

// Println prints given arguments with newline.
func Println(a ...interface{}) (n int, err error) {
	return defaultPrettyPrinter.Println(a...)
}

// Sprint formats given arguemnts and returns the result as string.
func Sprint(a ...interface{}) string {
	return defaultPrettyPrinter.Sprint(a...)
}

// Sprintf formats with pretty print and returns the result as string.
func Sprintf(format string, a ...interface{}) string {
	return defaultPrettyPrinter.Sprintf(format, a...)
}

// Sprintln formats given arguemnts with newline and returns the result as string.
func Sprintln(a ...interface{}) string {
	return defaultPrettyPrinter.Sprintln(a...)
}

// Fprint prints given arguments to a given writer.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return defaultPrettyPrinter.Fprint(w, a...)
}

// Fprintf prints format to a given writer.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return defaultPrettyPrinter.Fprintf(w, format, a...)
}

// Fprintln prints given arguments to a given writer with newline.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return defaultPrettyPrinter.Fprintln(w, a...)
}

// Errorf formats given arguments and returns it as error type.
func Errorf(format string, a ...interface{}) error {
	return defaultPrettyPrinter.Errorf(format, a...)
}

// Fatal prints given arguments and finishes execution with exit status 1.
func Fatal(a ...interface{}) {
	defaultPrettyPrinter.Fatal(a...)
}

// Fatalf prints a given format and finishes execution with exit status 1.
func Fatalf(format string, a ...interface{}) {
	defaultPrettyPrinter.Fatalf(format, a...)
}

// Fatalln prints given arguments with newline and finishes execution with exit status 1.
func Fatalln(a ...interface{}) {
	defaultPrettyPrinter.Fatalln(a...)
}

// Change Print* functions' output to a given writer.
// For example, you can limit output by ENV.
//
//	func init() {
//		if os.Getenv("DEBUG") == "" {
//			pp.SetDefaultOutput(ioutil.Discard)
//		}
//	}
func SetDefaultOutput(o io.Writer) {
	defaultPrettyPrinter.SetOutput(o)
}

// GetOutput returns pp's default output.
func GetDefaultOutput() io.Writer {
	return defaultPrettyPrinter.GetOutput()
}

// Change Print* functions' output to default one.
func ResetDefaultOutput() {
	defaultPrettyPrinter.ResetOutput()
}

// SetColorScheme takes a colorscheme used by all future Print calls.
func SetColorScheme(scheme ColorScheme) {
	defaultPrettyPrinter.SetColorScheme(scheme)
}

// ResetColorScheme resets colorscheme to default.
func ResetColorScheme() {
	defaultPrettyPrinter.ResetColorScheme()
}

// SetMaxDepth sets the printer's Depth, -1 prints all
func SetDefaultMaxDepth(v int) {
	defaultPrettyPrinter.maxDepth = v
}

// WithLineInfo add file name and line information to output
// call this function with care, because getting stack has performance penalty
var WithLineInfo bool
