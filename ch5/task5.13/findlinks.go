// Findlinks crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"net/url"

	"os"

	"path/filepath"
	"strings"

	"io"
	"net/http"

	"gopl.io/ch5/links"
)

const maxPageCount = 100

var (
	domain    string
	pageCount int
)

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
func crawl(rawurl string) []string {
	log.Printf("visit %s", rawurl)

	if err := save(rawurl); err != nil {
		log.Printf("save %s as file: %v", rawurl, err)
		// If an error occurs while saving, then do not follow extracted URLs ?
		return nil
	}

	list, err := links.Extract(rawurl)
	if err != nil {
		log.Printf("extract links from %s: %v", rawurl, err)
		return nil
	}

	return list
}

func save(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("parse %s to save as html: %v", rawurl, err)
	}
	// It's not an error. But if return nil then links will be extracted from this link?
	if u.Host != domain {
		log.Printf("%s hasn't valid domain", u.String())
		return nil
		//return fmt.Errorf("%s hasn't valid domain", u.String())
	}

	dir, file := buildPath(u.Path)
	if err := createDir(dir); err != nil {
		return err
	}
	if err := saveURL(rawurl, file); err != nil {
		return err
	}

	log.Printf("saved at %s", file)

	return nil
}

// createDir creates directory with all above directories if doesn't exist.
func createDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("create directory %s: %v", dir, err)
		}
	}

	log.Printf("create directory %s", dir)
	return nil
}

func saveURL(url, dest string) (err error) {
	if pageCount >= maxPageCount {
		// The crawl ends after it has visited all links in the queue.
		// It can work with error for a long time. Special type of error or Fatalf will be better decision.
		return fmt.Errorf("max page count %d was reached", maxPageCount)
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("get file on %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Open a file for writing or truncate the old.
	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("create %s: %v", dest, err)
	}
	// Are there any better way to close the file?
	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	// Use io.Copy to just dump the response body to the file.
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("copy %s to %s: %v", url, dest, err)
	}

	pageCount++
	return nil
}

func buildPath(urlPath string) (dir, file string) {
	linkPath := filepath.Join(domain, urlPath)

	if strings.HasSuffix(urlPath, "/") {
		return linkPath, filepath.Join(linkPath, "index.html")
	} else {
		dir, _ := filepath.Split(linkPath)
		return dir, linkPath
	}
}
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.

	// Site url should have trailing slash!
	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("parse url failed: %v", err)
	}
	domain = u.Host

	breadthFirst(crawl, []string{u.String()})
}
