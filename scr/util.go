package debug

import (
	"fmt"
	"time"
)

var Mode bool = true

func CurrentTime() string {
	hour := time.Now().Local().Hour()
	min := time.Now().Local().Minute()
	sec := time.Now().Local().Second()
	mili := time.Now().Local().Nanosecond() / 1e6
	return fmt.Sprintf("%s|%02d:%02d:%02d:%03d|%s", COLOR_GREEN, hour, min, sec, mili, COLOR_RESET)
}

func Print(color string, count int32, tYPE string, args interface{}) {
	fmt.Printf("%s %s{%03d} %s%s %s\n", CurrentTime(), color, count, tYPE, COLOR_RESET, args)
}
