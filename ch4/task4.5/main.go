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
		"m",
		"m",
	}

	//tc = unique(tc)
	tc = optimizedUnique(tc)
	//tc = optimizedUniqueWrong(tc)
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

func optimizedUniqueWrong(s []string) []string {
	duplicate := false
	i := 1
	for j := 0; i < len(s); i++ {
		if duplicate && s[i] != s[i-1] {
			copy(s[j+1:], s[i:])
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

func optimizedUnique(s []string) []string {
	writePtr := 1
	readPtr := 1
	for ; readPtr < len(s); readPtr++ {
		if s[readPtr] != s[readPtr-1] {
			s[writePtr] = s[readPtr]
			writePtr++
		}
	}
	return s[:writePtr]
}
