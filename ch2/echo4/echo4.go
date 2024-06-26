package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "раздедитьель")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
