package debug

import (
	"fmt"
)

const version string = "1.0.7"

func init() {
	if Mode {
		fmt.Println(CurrentTime(), "Currently using HaraldWik's debug system", version)
		println("\n#########  ########## #########  ###    ###  ########  \n###    ### ###        ###    ### ###    ### ###    ### \n###    ### ###        ###    ### ###    ### ###        \n###    ### ########   #########  ###    ### ###        \n###    ### ###        ###    ### ###    ### ###   #### \n###    ### ###        ###    ### ###    ### ###    ### \n#########  ########## #########   ########   ########  \n")
	}
}
