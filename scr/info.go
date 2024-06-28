package debug

import (
	"fmt"
	"strings"
)

const INFO string = COLOR_BLUE + "[INFO]" + COLOR_RESET

var InfoCount int32

func Info(text string) {
	if strings.TrimSpace(text) != "" {
		fmt.Println(INFO+COLOR_GRAY, InfoCount, COLOR_RESET+text)
		InfoCount++
	}
	if Mode && strings.TrimSpace(text) == "" {
		Warning("No info given!")
	}
}
