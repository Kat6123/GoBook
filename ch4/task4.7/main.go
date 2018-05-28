package main

import (
	"fmt"
)

func main() {

}


func reverse(b []byte) []byte {
	s := []rune(string(s))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return []byte(string(s))
}
