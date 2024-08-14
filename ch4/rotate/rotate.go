package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	step := 4

	fmt.Println(array)
	array = rotate(array, step)
	fmt.Println(array)
}

func rotate(slice []int, step int) []int {
	if step < 0 {
		step = len(slice) + step
		if step < 0 {
			step *= -1
		}
	}
	reverse(slice[:step])
	reverse(slice[step:])
	reverse(slice)
	return slice
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
