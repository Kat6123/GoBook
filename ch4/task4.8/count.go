// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func identifyRune(r rune, m map[string]int) {
	switch {
	case unicode.IsLetter(r):
		m["letter"]++
	case unicode.IsDigit(r):
		m["digit"]++
	case unicode.IsControl(r):
		m["control"]++
	case unicode.IsGraphic(r):
		m["graphic"]++
	case unicode.IsLower(r):
		m["lower"]++
	case unicode.IsMark(r):
		m["mark"]++
	case unicode.IsPunct(r):
		m["punct"]++
	default:
		m["other"]++
	}
}

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		fmt.Printf("Echo %q\n", r)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}

		identifyRune(r, counts)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
