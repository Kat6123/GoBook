// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("analyze %s as html: %v", url, err)
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

const alpha = "abcdefghijklmnopqrstuvwxyz"

func startElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		printComment(n)
	case html.ElementNode:
		openElement(n)
	case html.TextNode:
		printText(n)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		closeElement(n)
	}
}

func openElement(n *html.Node) {
	fmt.Printf("%*s<%s", depth*2, "", n.Data)
	printTags(n)

	if n.FirstChild == nil {
		fmt.Printf("/>\n")
	} else {
		fmt.Printf(">\n")
		depth++
	}
}

func printTags(n *html.Node) {
	for i := range n.Attr {
		attr := n.Attr[i]
		fmt.Printf(" %s=%q", attr.Key, attr.Val)
	}
}

func closeElement(n *html.Node) {
	if n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func printComment(n *html.Node) {
	fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
}

func printText(n *html.Node) {
	// Otherwise n.Data consists of \n and spaces.
	if strings.ContainsAny(strings.ToLower(n.Data), alpha) {
		n.Data = strings.Trim(n.Data, "\n")
		n.Data = strings.Replace(n.Data, "\n", fmt.Sprintf("\n%*s", depth*2, ""), -1)

		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

//!-startend
