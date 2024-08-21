package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error open file:\t%v", err)
		}

		input := bufio.NewScanner(file)
		input.Split(bufio.ScanWords)

		for input.Scan() {
			counts[onlyWord(input.Text())]++
		}

		fmt.Fprintf(os.Stdout, "From file:\t%s\nWord:\t\tCount:\n", filename)
		for word, count := range counts {
			fmt.Fprintf(os.Stdout, "%s\t\t%d\n", word, count)
		}

	}
}

func onlyWord(str string) string {
	newStr := strings.ReplaceAll(str, ".", "")
	newStr = strings.ReplaceAll(newStr, ",", "")
	newStr = strings.ReplaceAll(newStr, "!", "")
	newStr = strings.ReplaceAll(newStr, "?", "")
	newStr = strings.ReplaceAll(newStr, "“", "")
	newStr = strings.ReplaceAll(newStr, "”", "")
	return newStr
}
