// Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join.

// Results:
// executing: ./1.3 a line to be echoed for this test
// 10,000,000 times
// Test concat string: 7.18s elapsed
//   Test Join string: 2.21s elapsed

// HOWEVER
// executing without any input: ./1.3
// 1,000,000,000 times (100x previous test)
// Test concat string: 1.43s elapsed
//   Test Join string: 5.46s elapsed

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	testCount := 10000000
	// testCount := 1000000000 // without any command strings, 100 times previous

	start := time.Now()
	for i := 1; i < testCount; i++ {
		// stringConcat()
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
	fmt.Printf("Test concat string: %.2fs elapsed\n", time.Since(start).Seconds())

	start2 := time.Now()
	for i := 1; i < testCount; i++ {
		// stringJoin()
		strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("  Test Join string: %.2fs elapsed\n", time.Since(start2).Seconds())

}

// stringConcat produces a string that echos the command line
// using a for loop and += string concatenation
func stringConcat() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func stringJoin() string {
	return strings.Join(os.Args[1:], " ")
}
