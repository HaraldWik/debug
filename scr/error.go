package debug

var errorCount int32

func Error(err error) {
	if err != nil {
		Print(COLOR_RED, errorCount, "[ERROR]  ", err)
		errorCount++
	}
}
