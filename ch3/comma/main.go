package main

import "fmt"

func main() {
	var strNumInt string = "123455456"
	var strNumFloat string = "1234.55456"
	fmt.Println(comma(strNumInt))
	fmt.Println(commaNonRecursion(strNumInt))
	fmt.Println(commaFloat(strNumFloat))

	var str1 string = "table"
	var str2 string = "bleat"
	fmt.Printf("%v\n", isAnagram(str1, str2))
}
