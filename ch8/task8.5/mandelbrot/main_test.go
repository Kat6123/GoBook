package main

import "testing"

func BenchmarkConstructImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		constructImage()
	}
}

func BenchmarkConstructImageParallelY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		constructImageParallelY()
	}
}

func BenchmarkConstructImageParallelX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		constructImageParallelX()
	}
}

func BenchmarkConstructImageParallelYAndX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		constructImageParallelYAndX()
	}
}

func BenchmarkConstructImageBuffParallelY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		constructImageBuffParallelY()
	}
}
