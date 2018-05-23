package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("is Anagram: ")
	fmt.Println(isAnagram("Привет", "тевирП"))
	fmt.Println(isAnagram("Привет", "теПври"))
	fmt.Println(isAnagram("Привет", "Пока"))

	fmt.Println("is Reversed: ")
	fmt.Println(isReversed("Привет", "тевирП"))
	fmt.Println(isReversed("Привет", "теПври"))
	fmt.Println(isReversed("abra", "arba"))
}

// isAnagram check if strings consist of the same letters
func isAnagram(s1 string, s2 string) bool{
	for _, r := range s2{
		if !strings.ContainsRune(s1, r){
			return false
		}
	}
	return true
}

// isReversed check if second string is reversed version of first
func isReversed(s1 string, s2 string) bool{
	r1, r2 := []rune(s1), []rune(s2)
	l1, l2 := len(r1), len(r2)

	if l1 != l2{
		return false
	}

	for i := 0; i < l1 / 2; i++{
		if r1[i] != r2[l2 - 1 - i] || r2[i] != r1[l1 - 1 - i]{
			return false
		}
	}

	return true
}