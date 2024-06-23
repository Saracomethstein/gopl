package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	cycles := 5

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, cycles)
	})
	http.HandleFunc("/snippet", func(w http.ResponseWriter, r *http.Request) {
		cycles = showSnippet(w, r, "cycles")
		if cycles > 0 {
			lissajous(w, cycles)
		} else {
			http.NotFound(w, r)
		}
	})
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func showSnippet(w http.ResponseWriter, r *http.Request, snippetName string) (resultSnippet int) {
	snippet, err := strconv.Atoi(r.URL.Query().Get(snippetName))

	if err != nil || snippet < 1 {
		return
	}
	resultSnippet = snippet
	return resultSnippet
}
