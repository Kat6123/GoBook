package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type links map[string][]string

// visit appends to links each link found in n, and returns the result.
func visit(links links, n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			appendLink("a", "href", n, links)
		case "script":
			appendLink("script", "src", n, links)
		case "img":
			appendLink("img", "src", n, links)
		case "link":
			if isCSS(n) {
				appendLink("css", "href", n, links)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
}

func isCSS(n *html.Node) bool {
	for _, a := range n.Attr {
		if a.Key == "rel" && a.Val == "stylesheet" {
			return true
		}
	}
	return false
}

func appendLink(element, ref string, n *html.Node, links links) {
	for _, a := range n.Attr {
		if a.Key == ref {
			links[element] = append(links[element], a.Val)
		}
	}
}

//!+
func main() {
	for _, url := range os.Args[1:] {
		l := links{}
		if err := findLinks(l, url); err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for linkType := range l {
			for i := range l[linkType] {
				fmt.Printf("%s: %s\n", linkType, l[linkType][i])
			}
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(links links, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visit(links, doc)
	return nil
}

//!-
