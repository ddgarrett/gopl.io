// Exercise 1.1: Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.

// See page 8.

// Modify the original Echo3 program.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args, " "))
}
