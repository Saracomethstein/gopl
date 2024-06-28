package main

import (
	"popcount/popcount"
	"testing"
)

func BenchmarkPopCountOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountOld(123421341234)
	}
}

func BenchmarkPopCountNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountNew(123421341234)
	}
}
