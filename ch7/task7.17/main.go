package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"log"
)

func main() {
	f, err := os.Open("source.html")
	if err != nil {
		log.Fatalf("open file %s: %v", "source.html", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			// Fatalf in defer?
			log.Fatalf("close file: %v", err)
		}
	}()

	dec := xml.NewDecoder(f)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("xmlselect: %v\n", err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			//a := tok.Attr[0]
			//a.Name.Local == "b"
			//a.Value == ""

			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
