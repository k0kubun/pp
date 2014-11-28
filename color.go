package pp

import (
	"fmt"
	"strings"
)

var (
	codeByColor = map[string]int{
		"red":     31,
		"green":   32,
		"yellow":  33,
		"blue":    34,
		"magenta": 35,
		"cyan":    36,
	}
)

func color(text, color string) string {
	if code, ok := codeByColor[color]; ok {
		return fmt.Sprintf("\033[%dm%s\033[0m", code, text)
	} else if code, ok := codeByColor[strings.ToLower(color)]; ok {
		return fmt.Sprintf("\033[%dm\033[1m%s\033[0m", code, text)
	} else {
		return text
	}
}
