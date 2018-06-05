package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestPrintComment(t *testing.T) {
	tt := []struct {
		testName     string
		commentNode  *html.Node
		depth        int
		expectedHtml string
	}{
		{
			testName: "comment",
			commentNode: &html.Node{
				Type: html.CommentNode,
				Data: "Simple comment",
			},
			expectedHtml: "<!--Simple comment-->\n",
		},
		{
			testName: "innerComment",
			commentNode: &html.Node{
				Type: html.CommentNode,
				Data: "Simple comment",
			},
			depth:        3,
			expectedHtml: "      <!--Simple comment-->\n",
		},
		{
			testName: "multilineComment",
			commentNode: &html.Node{
				Type: html.CommentNode,
				Data: "\nSimple comment\nMultiline comment",
			},
			depth:        0,
			expectedHtml: "<!--\nSimple comment\nMultiline comment-->\n",
		},
		{
			testName: "innerMultilineComment",
			commentNode: &html.Node{
				Type: html.CommentNode,
				Data: "\nSimple comment\nMultiline comment",
			},
			depth:        3,
			expectedHtml: "      <!--\n      Simple comment\n      Multiline comment-->\n",
		},
		{
			testName: "emptyComment",
			commentNode: &html.Node{
				Type: html.CommentNode,
				Data: "",
			},
			depth:        0,
			expectedHtml: "<!---->\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			depth = tc.depth
			buffer := bytes.Buffer{}
			printComment(&buffer, tc.commentNode)
			assert.Equal(t, tc.expectedHtml, buffer.String())
		})
	}
}

func TestPrintText(t *testing.T) {
	tt := []struct {
		testName     string
		depth        int
		textNode     *html.Node
		expectedHtml string
	}{
		{
			"plain",
			0,
			&html.Node{
				Type: html.TextNode,
				Data: "Plain text"},
			"Plain text\n"},
		{
			"innerPlain",
			3,
			&html.Node{
				Type: html.TextNode,
				Data: "Plain text"},
			"      Plain text\n"},
		{
			"multilineText",
			0,
			&html.Node{
				Type: html.TextNode,
				Data: "Some text\nAgain\nAgain"},
			"Some text\nAgain\nAgain\n"},
		{
			"innerMultilineText",
			3,
			&html.Node{
				Type: html.TextNode,
				Data: "Some text\nAgain\nAgain"},
			"      Some text\n      Again\n      Again\n"},
		{
			// XXX?
			"multilineTextSurroundedWithEnters",
			3,
			&html.Node{
				Type: html.TextNode,
				Data: "\nSome text\nAgain\nAgain\n"},
			"      Some text\n      Again\n      Again\n"},
		{
			"notAlphabeticText",
			0,
			&html.Node{
				Type: html.TextNode,
				Data: "\n\n\n   \n"},
			""},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			buffer := bytes.Buffer{}
			depth = tc.depth
			printText(&buffer, tc.textNode)
			assert.Equal(t, tc.expectedHtml, buffer.String())
		})
	}
}

func TestPrintTags(t *testing.T) {
	tt := []struct {
		testName    string
		node        *html.Node
		expectedStr string
	}{
		{
			"tag",
			&html.Node{
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{Key: "href", Val: "/src/org"},
					{Key: "type", Val: ""},
					{Key: "href", Val: "org"}},
			},
			" href=\"/src/org\" type href=\"org\""},
		{
			"withoutTags",
			&html.Node{
				Type: html.ElementNode,
			},
			""},
		{
			"tagWithoutValue",
			&html.Node{
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{Key: "href"}},
			},
			" href"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			buffer := bytes.Buffer{}
			printTags(&buffer, tc.node)
			assert.Equal(t, tc.expectedStr, buffer.String())
		})
	}
}

func TestOpenElement(t *testing.T) {

}
