package debug

var errorCount int32

func Error(err error) {
	if err != nil && Mode {
		Print(COLOR_RED, errorCount, "[ERROR]  ", err)
		errorCount++
	}
}
