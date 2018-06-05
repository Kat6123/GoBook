package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Привет"
	fmt.Println(string(reverse([]byte(s))))
	fmt.Println(string(reverseV2([]byte(s))))
}

func reverse(b []byte) []byte {
	s := []rune(string(b)) // Allocation?

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return []byte(string(s))
}

func reverseV2(b []byte) []byte {
	bPtr := 0
	revPtr := len(b)
	rev := make([]byte, len(b))

	for _, i := utf8.DecodeRune(b); bPtr < len(b); {
		copy(rev[revPtr-i:revPtr], b[bPtr:bPtr+i])
		bPtr += i
		revPtr -= i
	}

	return rev
}
