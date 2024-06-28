package debug

import (
	"fmt"
	"strings"
)

const WARNING string = COLOR_ORANGE + "[WARNING]" + COLOR_RESET

var WarningCount int32

func Warning(text string) {
	if strings.TrimSpace(text) != "" {
		fmt.Println(WARNING+COLOR_GRAY, WarningCount, COLOR_RESET+text)
		WarningCount++
	}
	if Mode && strings.TrimSpace(text) == "" {
		Warning("No warning given!")
	}
}
