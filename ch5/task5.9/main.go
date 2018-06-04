package main

import (
	"fmt"
	"strings"
)

func main() {
	f := "abcdefg$fooasmmsa$foo"
	fmt.Println(expand(f, replacer))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func replacer(s string) string {
	return s + "\n"
}
