// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		// Check if url has "http://" prefix
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		// Get
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s\n", err)
			return
		}
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %s\n", url, err)
			return
		}
	}
}

//!-
