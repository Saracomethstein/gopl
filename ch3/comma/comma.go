package main

import "fmt"

func main() {
	var str string = "1234511121"
	fmt.Printf("%s\n", comma(str))
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
