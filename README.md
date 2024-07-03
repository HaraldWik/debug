# Debug
---
**Description**
---
Makes debuging easier and prettier!

---
**Exampels**
---
You can use the "_debug.Error()_" to handle errors instead of doing "_if err != nil_"
```golang
package main

import (
	"errors"

	debug "github.com/HaraldWik/debug/scr"
)

func FunctionThatReturnsError() error {
	return errors.New("code is too bad")
}

func main() {
	debug.Error(FunctionThatReturnsError())
}
```
---
You can also use it to display information, so the user can know whats happening!
```golang
package main

import (
	"fmt"

	debug "github.com/HaraldWik/debug/scr"
)

func main() {
	debug.Info("Program started!")
	fmt.Println("Hello, World!")
	debug.Info("Printed 'Hello, World!")
	fmt.Println("Bye, World!")
	debug.Info("Printed 'Bye, World!'")
	fmt.Print("\n")
	debug.Info("Printed new line")
	debug.Info("Program closed!")
}
```
---
**Visuals***
![image](https://github.com/HaraldWik/debug/assets/167028112/68819d10-19be-4659-a55e-df4b74c4ba91)
---
