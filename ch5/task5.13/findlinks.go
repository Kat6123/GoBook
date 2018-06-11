// Findlinks crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

const maxSeenCount = 100

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
			if len(seen) >= maxSeenCount {
				break
			}
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(rawurl string) []string {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Printf("parse %s has failed: %v", rawurl, err)
		return nil
	}

	if u.Host != domain {
		log.Printf("%s hasn't valid domain", u.String())
		return nil
	}

	log.Printf("visit %s", rawurl)

	if err := save(u); err != nil {
		log.Printf("save %s as file: %v", rawurl, err)
		return nil
	}

	list, err := links.Extract(rawurl)
	if err != nil {
		log.Printf("extract links from %s: %v", rawurl, err)
		return nil
	}

	return list
}

func save(u *url.URL) error {
	dir, file := buildPath(u.Path)
	if err := createDir(dir); err != nil {
		return err
	}
	if err := saveURL(u.String(), file); err != nil {
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
		log.Printf("create directory %s", dir)
	}

	return nil
}

func saveURL(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("get file on %s: %v", url, err)
	}
	defer resp.Body.Close()

	if err := write(resp.Body, dest); err != nil {
		return fmt.Errorf("write to file: %v", err)
	}

	pageCount++
	return nil
}

func write(r io.ReadCloser, dest string) (err error) {
	// Open a file for writing or truncate the old.
	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("create %s: %v", dest, err)
	}

	// Use io.Copy to just dump the response body to the file.
	_, err = io.Copy(file, r)
	if err != nil {
		err = fmt.Errorf("copy response to %s: %v", dest, err)
	}
	if closeErr := file.Close(); err == nil && closeErr != nil {
		err = fmt.Errorf("close file %q: %v", dest, closeErr)
	}
	return
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
