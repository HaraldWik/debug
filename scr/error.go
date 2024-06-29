package debug

import (
	"fmt"
	"time"
)

const ERROR string = COLOR_RED + "[ERROR]" + COLOR_RESET

var ErrorCount int32

func Error(err error) {
	if err != nil {
		fmt.Printf("%s %s%ds-%dm-%dh %d%s %s\n", ERROR, COLOR_GRAY, time.Now().Second(), time.Now().Minute(), time.Now().Hour(), InfoCount, COLOR_RESET, err)
		ErrorCount++
	}
}
