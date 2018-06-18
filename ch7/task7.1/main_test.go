package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter_Write(t *testing.T) {
	tt := []struct {
		testName      string
		input         string
		expectedCount wordCounter
	}{
		{
			testName:      "multiword",
			input:         "Hi there it's new string  ",
			expectedCount: 5,
		},
		{
			testName:      "empty",
			input:         "",
			expectedCount: 0,
		},
		{
			testName:      "spaces",
			input:         "      ",
			expectedCount: 0,
		},
		{
			testName:      "utf-8",
			input:         "Прив ет",
			expectedCount: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			var counter wordCounter
			counter.Write([]byte(tc.input))

			assert.Equal(t, tc.expectedCount, counter)
		})
	}
}

func TestLineCounter_Write(t *testing.T) {
	tt := []struct {
		testName      string
		input         string
		expectedCount lineCounter
	}{
		{
			testName:      "multiline",
			input:         "Hi there\n it's new string\n  ",
			expectedCount: 3,
		},
		{
			testName:      "empty",
			input:         "",
			expectedCount: 0,
		},
		{
			testName:      "enters",
			input:         "\n\n\n",
			expectedCount: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			var counter lineCounter
			counter.Write([]byte(tc.input))

			assert.Equal(t, tc.expectedCount, counter)
		})
	}
}
