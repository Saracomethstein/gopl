package main

import (
	mondelbrot "ch3/mondelbrot"
	newton "ch3/newton"
	surface "ch3/surface"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/surface", func(w http.ResponseWriter, r *http.Request) {
		surface.Surface(w, 100)
	})

	http.HandleFunc("/mondelbrot", func(w http.ResponseWriter, r *http.Request) {
		mondelbrot.Mondelbrot(w, r, 0, 0)
	})

	http.HandleFunc("/newton", func(w http.ResponseWriter, r *http.Request) {
		newton.NewtonMondelbrot(w, r)
	})

	http.HandleFunc("/newton64", func(w http.ResponseWriter, r *http.Request) {
		newton.NewtonComplex64Handler(w, r)
	})

	http.HandleFunc("/newton128", func(w http.ResponseWriter, r *http.Request) {
		newton.NewtonComplex128Handler(w, r)
	})

	http.HandleFunc("/surface/snippet", func(w http.ResponseWriter, r *http.Request) {
		params := getQueryParams(r)
		cells, _ := strconv.Atoi(params["cells"])
		if cells == 0 {
			cells = 100
		}
		surface.Surface(w, cells)
	})

	http.HandleFunc("/mondelbrot/snippet", func(w http.ResponseWriter, r *http.Request) {
		params := getQueryParams(r)
		x, _ := strconv.ParseFloat(params["x"], 64)
		y, _ := strconv.ParseFloat(params["y"], 64)
		mondelbrot.Mondelbrot(w, r, x, y)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getQueryParams(r *http.Request) map[string]string {
	params := make(map[string]string)
	query := r.URL.Query()
	for key, values := range query {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}
	return params
}
