package debug

import "fmt"

const ERROR string = COLOR_RED + "[ERROR]" + COLOR_RESET

var ErrorCount int32

func Error(err error) {
	if err != nil {
		fmt.Println(ERROR+COLOR_GRAY, ErrorCount, COLOR_RESET, err)
		ErrorCount++
	}
	if err == nil {
		Warning("No error given!")
	}
}
