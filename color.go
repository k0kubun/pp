package pp

import (
	"bytes"
	"reflect"
	"sort"
	"strconv"
)

var (
	colorByFlag = map[int]FlagSet{
		30: Black,
		31: Red,
		32: Green,
		33: Yellow,
		34: Blue,
		35: Magenta,
		36: Cyan,
		37: White,
		40: BackBlack,
		41: BackRed,
		42: BackGreen,
		43: BackYellow,
		44: BackBlue,
		45: BackMagenta,
		46: BackCyan,
		47: BackWhite,
		1:  Bold,
	}

	// Used for keeping track of all flags sorted by their colorCodes
	// Useful for testing
	flags []int

	defaultScheme = ColorScheme{
		Bool:            Cyan | Bold,
		Integer:         Blue | Bold,
		Float:           Magenta | Bold,
		String:          Red,
		StringQuotation: Red | Bold,
		FieldName:       Yellow,
		PointerAdress:   Blue | Bold,
		Nil:             Cyan | Bold,
		Time:            Blue | Bold,
		StructName:      Green,
		ObjectLength:    Blue,
	}
)

type FlagSet uint

const (
	Black FlagSet = 1 << iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BackBlack
	BackRed
	BackGreen
	BackYellow
	BackBlue
	BackMagenta
	BackCyan
	BackWhite
	Bold
)

type ColorScheme struct {
	Bool            FlagSet
	Integer         FlagSet
	Float           FlagSet
	String          FlagSet
	StringQuotation FlagSet
	FieldName       FlagSet
	PointerAdress   FlagSet
	Nil             FlagSet
	Time            FlagSet
	StructName      FlagSet
	ObjectLength    FlagSet
}

func init() {
	for key, _ := range colorByFlag {
		flags = append(flags, key)
	}
	sort.Ints(flags)
}

// Makes sure each Color is set, if not reverts that field to default value
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

func colorize(text string, color FlagSet) string {
	buf := bytes.NewBufferString("")

	for _, id := range flags {
		flag := colorByFlag[id]
		if flag&color != 0 {
			buf.WriteString("\x1b[")
			buf.WriteString(strconv.Itoa(id))
			buf.WriteString("m")
		}
	}
	buf.WriteString(text)
	buf.WriteString("\x1b[0m")

	return buf.String()
}
