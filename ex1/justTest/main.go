package main

import (
	"fmt"
	"os"
)

func printf(strFmt string, a ...interface{}) {
	fmt.Printf(strFmt, a...)
}

func main() {
	printf("%T \n", os.Args)

	fmt.Printf("%T \n", os.Args)

	printf("%v \n", os.Args)

	fmt.Printf("%v \n", os.Args)
}
