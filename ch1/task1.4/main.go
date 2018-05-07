// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type fileApp map[string]map[string]int

func main() {
	counts := make(fileApp)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "StdIn")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %s\n", err)
				continue
			}
			defer f.Close()
			countLines(f, counts, arg)
		}
	}
	for line, input := range counts {
		fmt.Println(line)
		for fileName, count := range input {
			if count > 1 {
				fmt.Printf("\t%s\t%d\n", fileName, count)
			}
		}
	}
}

func countLines(f *os.File, counts fileApp, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		_, ok := counts[input.Text()]
		if ok != true {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][fileName]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
