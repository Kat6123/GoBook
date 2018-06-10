package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildPath(t *testing.T) {
	tt := []struct {
		testName                  string
		url                       string
		domain                    string
		dirExpected, fileExpected string
	}{
		{
			testName:     "path with document",
			url:          "/path/url/some.html",
			domain:       "gopl.io",
			dirExpected:  "gopl.io/path/url/",
			fileExpected: "gopl.io/path/url/some.html",
		},
		{
			testName:     "path ends with slash",
			url:          "/path/url/",
			domain:       "gopl.io",
			dirExpected:  "gopl.io/path/url",
			fileExpected: "gopl.io/path/url/index.html",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			domain = tc.domain
			dir, file := buildPath(tc.url)

			assert.Equal(t, tc.dirExpected, dir, "directory path not equel")
			assert.Equal(t, tc.fileExpected, file, "file path not equel")
		})
	}
}
