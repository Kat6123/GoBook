package main

import "fmt"

func main() {
	tc := []string{
		"abc",
		"abc",
		"abc",
		"d",
		"b",
		"b",
		"ad",
		"ad",
		"ad",
	}

	//tc = unique(tc)
	//tc = optimizedUnique(tc)
	fmt.Println(tc)
}

func unique(s []string) []string {
	i := 1
	for i < len(s) {
		if s[i] == s[i-1] {
			copy(s[i-1:], s[i:])
			s = s[:len(s)-1]
			continue
		}
		i++
	}

	return s[:i]
}

func optimizedUnique(s []string) []string {
	duplicate := false
	i := 1
	for j := 0; i < len(s); i++ {
		if duplicate && s[i] != s[i-1] {
			copy(s[j + 1:], s[i:])
			s = s[:len(s)-(i-j)]
			continue
		}
		if !duplicate && s[i] == s[i-1] {
			duplicate = true
			j = i - 1
		}
	}

	return s[:i]
}
