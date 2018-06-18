// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Len counts elements in the set.
func (s *IntSet) Len() int {
	count := 0

	for i := range s.words {
		if s.words[i] == 0 {
			continue
		}

		for j := uint(0); j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				count++
			}
		}
	}

	return count
}

// Remove removes the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)

	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Clear clears the set.
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy creates copy of the set and returns set pointer.
func (s *IntSet) Copy() *IntSet {
	newSet := IntSet{
		words: make([]uint64, len(s.words)),
	}

	for i := range s.words {
		newSet.words[i] = s.words[i]
	}

	return &newSet
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			s.words[i] = 0
		} else {
			s.words[i] &= t.words[i]
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
}

// SymmetricDifferenceWith sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, t.words[i])
		} else {
			s.words[i] ^= t.words[i]
		}
	}
}