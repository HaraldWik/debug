package debug

import (
	"fmt"
	"time"
)

const version string = "1.0.3"

func init() {
	if Mode {
		fmt.Println(time.Now().Second(), time.Now().Minute(), time.Now().Hour(), time.Now().Day(), "Currently using HaraldWik's debug system", version)
	}
}
