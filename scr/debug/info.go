package debug

import (
	"fmt"
	"strings"
	"time"
)

const INFO string = COLOR_BLUE + "[INFO]" + COLOR_RESET

var InfoCount int32

func Info(text string) {
	if strings.TrimSpace(text) != "" {
		fmt.Printf("%s %s%ds-%dm-%dh %d%s %s\n", INFO, COLOR_GRAY, time.Now().Second(), time.Now().Minute(), time.Now().Hour(), InfoCount, COLOR_RESET, text)
		InfoCount++
	}
	if Mode && strings.TrimSpace(text) == "" {
		Warning("No info given!")
	}
}
