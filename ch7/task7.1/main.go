package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type (
	wordCounter int
	lineCounter int
)

func (c *wordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		// Should I add info about p?
		// How return count of bytes which was read?
		return 0, fmt.Errorf("count words in %q with scanner: %v", string(p), err)
	}

	return len(p), nil
}

func (c *lineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("count lines in %q with scanner: %v", string(p), err)
	}

	return len(p), nil
}

func main() {
	var (
		w wordCounter
		l lineCounter
	)
	fmt.Fprintf(&w, " Somew string привет")
	fmt.Println(w)

	fmt.Fprintf(&l, " Somew string\n New line\n Another one")
	fmt.Println(l)
}
