package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// fetch выводит ответ на запрос по заданному URL
func fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(addPrefix(url))

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reader %s: %b\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", b)
	}
}

func addPrefix(url string) (result string) {
	var prefix string = "http://"

	if url[:7] != "http://" {
		result = prefix + url
	} else {
		return url
	}

	return result
}
