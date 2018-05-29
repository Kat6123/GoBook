// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const arrLen = 6

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"

	//Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints [arrLen]int
		i := 0
		for _, s := range strings.Fields(input.Text()) {
			if i >= arrLen {
				break
			}
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints[i] = int(x)
			i++
		}
		reverse(&ints)
		fmt.Printf("%v\n", ints)
	}
	//NOTE: ignoring potential errors from input.Err()
}

// I can access elements without (*s)[i]
func reverse(s *[arrLen]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
