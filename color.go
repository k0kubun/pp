package pp

import (
	"fmt"
	"reflect"
)

const (
	// NoColor no color
	NoColor uint16 = 1 << 15
)

const (
	// Foreground colors for ColorScheme.
	_ uint16 = iota | NoColor
	// Black color black
	Black
	// Red color red
	Red
	// Green color green
	Green
	// Yellow color yellow
	Yellow
	// Blue color blue
	Blue
	// Magenta color magenta
	Magenta
	// Cyan color cyan
	Cyan
	// White color white
	White
	bitsForeground       = 0
	maskForeground       = 0xf
	ansiForegroundOffset = 30 - 1
)

const (
	// Background colors for ColorScheme.
	_ uint16 = iota<<bitsBackground | NoColor
	// BackgroundBlack background color black
	BackgroundBlack
	// BackgroundRed background color red
	BackgroundRed
	// BackgroundGreen background color green
	BackgroundGreen
	// BackgroundYellow background color yellow
	BackgroundYellow
	// BackgroundBlue background color blue
	BackgroundBlue
	// BackgroundMagenta background color magenta
	BackgroundMagenta
	// BackgroundCyan background color cyan
	BackgroundCyan
	// BackgroundWhite background color white
	BackgroundWhite
	bitsBackground       = 4
	maskBackground       = 0xf << bitsBackground
	ansiBackgroundOffset = 40 - 1
)

const (
	// Bold flag for ColorScheme.
	Bold     uint16 = 1<<bitsBold | NoColor
	bitsBold        = 8
	maskBold        = 1 << bitsBold
	ansiBold        = 1
)

// ColorScheme To use with SetColorScheme.
type ColorScheme struct {
	Bool            uint16
	Integer         uint16
	Float           uint16
	String          uint16
	StringQuotation uint16
	EscapedChar     uint16
	FieldName       uint16
	PointerAdress   uint16
	Nil             uint16
	Time            uint16
	StructName      uint16
	ObjectLength    uint16
}

var (
	// ColoringEnabled DEPRECATED: Use PrettyPrinter.SetColoringEnabled().
	ColoringEnabled = true

	defaultScheme = ColorScheme{
		Bool:            Cyan | Bold,
		Integer:         Blue | Bold,
		Float:           Magenta | Bold,
		String:          Red,
		StringQuotation: Red | Bold,
		EscapedChar:     Magenta | Bold,
		FieldName:       Yellow,
		PointerAdress:   Blue | Bold,
		Nil:             Cyan | Bold,
		Time:            Blue | Bold,
		StructName:      Green,
		ObjectLength:    Blue,
	}
)

func (cs *ColorScheme) fixColors() {
	typ := reflect.Indirect(reflect.ValueOf(cs))
	defaultType := reflect.ValueOf(defaultScheme)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Uint() == 0 {
			field.SetUint(defaultType.Field(i).Uint())
		}
	}
}

func colorizeText(text string, color uint16) string {
	foreground := color & maskForeground >> bitsForeground
	background := color & maskBackground >> bitsBackground
	bold := color & maskBold

	if foreground == 0 && background == 0 && bold == 0 {
		return text
	}

	modBold := ""
	modForeground := ""
	modBackground := ""

	if bold > 0 {
		modBold = "\033[1m"
	}
	if foreground > 0 {
		modForeground = fmt.Sprintf("\033[%dm", foreground+ansiForegroundOffset)
	}
	if background > 0 {
		modBackground = fmt.Sprintf("\033[%dm", background+ansiBackgroundOffset)
	}

	return fmt.Sprintf("%s%s%s%s\033[0m", modForeground, modBackground, modBold, text)
}
