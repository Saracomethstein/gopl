package main

import (
	mondelbrot "ch3/mondelbrot"
	surface "ch3/surface"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/surface", func(w http.ResponseWriter, r *http.Request) {
		surface.Surface(w)
	})

	http.HandleFunc("/mondelbrot", func(w http.ResponseWriter, r *http.Request) {
		mondelbrot.Mondelbrot(w, r)
	})

	http.HandleFunc("surface/snippet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", surface.Cells)
		surface.Cells = showSnippet(r, "cells")
		if surface.Cells > 0 {
			surface.Surface(w)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func showSnippet(r *http.Request, snippetName string) (resultSnippet int) {
	snippet, err := strconv.Atoi(r.URL.Query().Get(snippetName))

	if err != nil || snippet < 1 {
		return
	}
	resultSnippet = snippet
	return resultSnippet
}
