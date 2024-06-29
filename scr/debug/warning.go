package debug

import (
	"fmt"
	"strings"
	"time"
)

const WARNING string = COLOR_ORANGE + "[WARNING]" + COLOR_RESET

var WarningCount int32

func Warning(text string) {
	if strings.TrimSpace(text) != "" {
		fmt.Printf("%s %s%ds-%dm-%dh %d%s %s\n", WARNING, COLOR_GRAY, time.Now().Second(), time.Now().Minute(), time.Now().Hour(), InfoCount, COLOR_RESET, text)
		WarningCount++
	}
	if Mode && strings.TrimSpace(text) == "" {
		Warning("No warning given!")
	}
}
