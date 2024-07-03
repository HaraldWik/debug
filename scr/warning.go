package debug

import (
	"strings"
)

var warningCount int32

func Warning(text string) {
	if strings.TrimSpace(text) != "" && Mode {
		Print(COLOR_ORANGE, warningCount, "[WARNING]", text)
		warningCount++
	}
	if Mode && strings.TrimSpace(text) == "" {
		Warning("No warning given!")
	}

}
