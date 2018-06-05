// Findlinks crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"net/url"

	"os"

	"path/filepath"

	"gopl.io/ch5/links"
)

const (
	root = "/var/crawler"
)

var domain string

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// Do I need to make abstraction under crawl? and pass save as anon func?
// TODO: add error handling
func crawl(rawurl string) []string {
	fmt.Println(rawurl)
	save(rawurl)

	list, err := links.Extract(rawurl)
	if err != nil {
		log.Print(err)
	}

	return list
}

func save(rawurl string) {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Printf("parse %s to save as html: %v", rawurl, err)
		return
	}
	if u.Host != domain {
		log.Printf("%s hasn't valid domain", u.String())
		return
	}

	linkPath := filepath.Join(root, domain, u.Path)
	if _, err := os.Stat(linkPath); os.IsNotExist(err) {
		os.MkdirAll(linkPath, os.ModePerm)
		log.Printf("create directory %s", linkPath)
	}

	//if strings.HasSuffix(u.Path, "/") {
	//	linkPath = filepath.Join(linkPath, "index.html")
	//}
	//log.Printf("saved at %s", linkPath)
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("parse url failed: %v", err)
	}
	domain = u.Host

	breadthFirst(crawl, []string{u.String()})
}
