package main

import (
	"log"
	"net/http"
	"strconv"
	surfacepac "surface/surface"
)

func main() {
	http.HandleFunc("/surface", func(w http.ResponseWriter, r *http.Request) {
		surfacepac.Surface(w)
	})

	// http.HandleFunc("/mondelbrot", func(w http.ResponseWriter, r *http.Request) {
	// 	mondelbrot.Mondelbrot()
	// })

	http.HandleFunc("surface/snippet", func(w http.ResponseWriter, r *http.Request) {
		surfacepac.Cells = showSnippet(r, "cells")
		if surfacepac.Cells > 0 {
			surfacepac.Surface(w)
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
