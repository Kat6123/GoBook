// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"strings"

	"io"
	"log"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("failed to load %s: %v", url, err)
		}
		defer resp.Body.Close()

		err = outline(url, resp.Body, os.Stdout)
		if err != nil {
			log.Printf("failed to outline html from %s: %v", url, err)
			continue
		}
	}
}

func outline(url string, reader io.Reader, w io.Writer) error {
	doc, err := html.Parse(reader)
	if err != nil {
		return fmt.Errorf("analyze %s as html: %v", url, err)
	}

	forEachNode(doc, startElement(w), endElement(w))

	return nil
}

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

var depth int

const alpha = "abcdefghijklmnopqrstuvwxyz"

func startElement(w io.Writer) func(n *html.Node) {
	return func(n *html.Node) {
		switch n.Type {
		case html.CommentNode:
			printComment(w, n)
		case html.ElementNode:
			openElement(w, n)
		case html.TextNode:
			printText(w, n)
		}
	}
}

func endElement(w io.Writer) func(n *html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode {
			closeElement(w, n)
		}
	}
}

func openElement(w io.Writer, n *html.Node) {
	fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
	printTags(w, n)

	if n.FirstChild == nil {
		fmt.Fprintf(w, "/>\n")
	} else {
		fmt.Fprintf(w, ">\n")
		depth++
	}
}

func printTags(w io.Writer, n *html.Node) {
	for i := range n.Attr {
		attr := n.Attr[i]
		if attr.Val != "" {
			fmt.Fprintf(w, " %s=%q", attr.Key, attr.Val)
		} else {
			fmt.Fprintf(w, " %s", attr.Key)
		}
	}
}

func closeElement(w io.Writer, n *html.Node) {
	if n.FirstChild != nil {
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}

func printComment(w io.Writer, n *html.Node) {
	n.Data = strings.Replace(n.Data, "\n", fmt.Sprintf("\n%*s", depth*2, ""), -1)
	fmt.Fprintf(w, "%*s<!--%s-->\n", depth*2, "", n.Data)
}

func printText(w io.Writer, n *html.Node) {
	// Otherwise n.Data consists of \n and spaces.
	if strings.ContainsAny(strings.ToLower(n.Data), alpha) {
		n.Data = strings.Trim(n.Data, "\n")
		n.Data = strings.Replace(n.Data, "\n", fmt.Sprintf("\n%*s", depth*2, ""), -1)
		fmt.Fprintf(w, "%*s%s\n", depth*2, "", n.Data)
	}
}
