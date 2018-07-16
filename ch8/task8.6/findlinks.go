// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"log"
	"os"

	"fmt"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	//fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

const d = 2

func main() {
	type item struct {
		l     []string
		depth int
	}

	worklist := make(chan item)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- item{os.Args[1:], 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		fmt.Println(n)
		it := <-worklist
		for _, link := range it.l {
			if !seen[link] {
				seen[link] = true
				fmt.Println(link, it.depth)
				if it.depth >= d {
					fmt.Println("countinue")
					continue
				}
				n++
				go func(link string, depth int) {
					worklist <- item{crawl(link), depth + 1}
				}(link, it.depth)
			} else {
				fmt.Println("dropped", link)
			}
		}
	}
}

//!-
