// Report //
/*
goos: darwin
goarch: amd64
pkg: popcount
cpu: Intel(R) Core(TM) i5-8500 CPU @ 3.00GHz
BenchmarkPopCountOld-6          1000000000               0.2558 ns/op
BenchmarkPopCountNew-6          287585209                4.089 ns/op
BenchmarkPopCount64-6           69005730                17.44 ns/op
BenchmarkPopCountDrop64-6       69414332                18.21 ns/op
PASS
ok      popcount        4.420s
*/
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

func BenchmarkPopCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount64(123421341234)
	}
}

func BenchmarkPopCountDrop64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountDrop64(123421341234)
	}
}
