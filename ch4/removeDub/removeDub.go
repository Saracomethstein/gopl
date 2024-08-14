package main

import (
	"fmt"
)

func main() {
	str := []string{"dummy", "one", "one", "two", "two", "potato", "potato", "dummy"}
	removeDup(&str)
	fmt.Println(str)
}

func removeDup(str *[]string) {
	count := 0
	for i := range *str {
		if i+1 > len(*str)-1 || (*str)[i] != (*str)[i+1] {
			(*str)[count] = (*str)[i]
			count++
		}
	}
}
