/*
MF конвертирует числовой аргумент в длину в Метрах и Футах.
*/
package main

import (
	"fmt"
	"lenghtconv/lenghtconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := lenghtconv.Foot(l)
		m := lenghtconv.Meter(l)

		fmt.Printf("%s = %s, %s = %s\n", f, lenghtconv.FToM(f), m, lenghtconv.MToF(m))
	}
}
