// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopl.io/ch4/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestTree_String(t *testing.T) {
	tt := []struct {
		testName       string
		treeBase       []int
		expectedString string
	}{
		{
			"tree",
			[]int{3, 2, 1, 4, -5, 5, 0, 1, 1},
			"{ -5 0 1 1 1 2 3 4 5 }",
		},
		{
			"empty",
			[]int{},
			"{ }",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			var root *tree
			for _, v := range tc.treeBase {
				root = add(root, v)
			}

			assert.Equal(t, tc.expectedString, root.String())
		})
	}
}
