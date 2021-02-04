// Modify dup2 to print the names of all files
// in which each duplicated line occurs.

// NOTE: would be more efficient to replace the map of maps
// "fileNames" with a fixed length array or slice that can be extended.
// Probably fixed length array would be enough since we don't expect
// too many file names to be specified

// ALSO note that in using os.Stdin for input requires a
// ctl-d to generate an EOF

// To run test use > ./1.4_dup2 f1.txt f2.txt

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1/1.4: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)

			inFiles := fileNames[line]
			// fmt.Printf("%v", inFiles)
			for fileName, fileN := range inFiles {
				fmt.Printf("\t%d\t%s\n", fileN, fileName)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++

		if fileNames[line] == nil {
			fileNames[line] = make(map[string]int)
		}

		fileNames[line][fileName]++

	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
