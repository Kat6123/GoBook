package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitedReader_Read(t *testing.T) {
	tt := []struct {
		testName       string
		max            int64
		input          string
		expectedOutput string
	}{
		{
			testName:       "limit more than stream length",
			max:            8,
			input:          "Hello",
			expectedOutput: "Hello",
		},
		{
			testName:       "limit equal to stream length",
			max:            5,
			input:          "Hello",
			expectedOutput: "Hello",
		},
		{
			testName:       "limit less than stream length",
			max:            2,
			input:          "Hello",
			expectedOutput: "He",
		},
		{
			testName:       "zero limit",
			max:            0,
			input:          "Hello",
			expectedOutput: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			lr := LimitReader(strings.NewReader(tc.input), tc.max)
			buf := new(bytes.Buffer)
			io.Copy(buf, lr)

			assert.Equal(t, tc.expectedOutput, buf.String())
		})
	}
}
