package main

import (
	"testing"
)

var tc = []string{
	"abc",
	"abc",
	"abc",
	"d",
	"b",
	"b",
	"ad",
	"ad",
	"ad",
}

func BenchmarkUnique(b *testing.B) {
	tcc := make([]string, len(tc))
	for n := 0; n < b.N; n++ {
		copy(tcc, tc)
		unique(tcc)
	}
}

func BenchmarkOptimizedUnique(b *testing.B) {
	tcc := make([]string, len(tc))
	for n := 0; n < b.N; n++ {
		copy(tcc, tc)
		optimizedUnique(tcc)
	}
}

func BenchmarkOptimizedUniqueWrong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		optimizedUniqueWrong(tc)
	}
}
