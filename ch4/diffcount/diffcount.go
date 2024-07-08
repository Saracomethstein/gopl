package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("diff:\t%d\n", diffCount(c1, c2))
}

func diffCount(x, y [32]byte) int {
	var count int

	for i := range x {
		if x[i] != y[i] {
			count++
		}
	}

	return count
}
