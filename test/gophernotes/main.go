package main

import (
	"fmt"
	"reflect"

	"github.com/cosmos72/gomacro/fast"
)

func RunGomacro(toeval string) []reflect.Value {
	interp := fast.New()
	vals, _ := interp.Eval(toeval)
	// NOT for simplicity, only use the first returned value
	return vals.ReflectValue()
}
func main() {
	fmt.Println(RunGomacro("1+1"))
}
