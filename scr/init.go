package debug

import "fmt"

const version string = "1.0.0"

func init() {
	if Mode {
		fmt.Printf("Currently using HaraldWik's debug system %s\n", version)
	}
}
