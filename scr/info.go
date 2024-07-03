package debug

import (
	"strings"
)

var infoCount int32

func Info(text string) {
	if strings.TrimSpace(text) != "" && Mode {
		Print(COLOR_LIGHT_GREEN, infoCount, "[INFO]   ", text)
		infoCount++
	}
	if Mode && strings.TrimSpace(text) == "" {
		Warning("No info given!")
	}
}
