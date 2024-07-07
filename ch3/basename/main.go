package main

import "fmt"

func main() {
	var path_1 string = "/home/sara/Documents/sources/gopl/ch3/basename/basename.go"
	var path_2 string = "/home/sara/Documents/sources/gopl/ch3/basename/basename.go"
	fmt.Printf("Basename (var 1.0)\n %v\n", basename(path_1))
	fmt.Printf("\nBasename (var 2.0)\n %v\n", basename2(path_2))
}
