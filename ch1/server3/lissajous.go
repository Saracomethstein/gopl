package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.RGBA{1, 1, 1, 1}, color.RGBA{0, 255, 0, 1}, color.RGBA{255, 0, 0, 1}, color.RGBA{0, 0, 255, 1}, color.RGBA{0, 0, 0, 0}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	blueIndex  = 3
	whiteIndex = 4
)

/*
Build command:

	go build main.go lissajous.go
	./main > name.gif
*/
func lissajous(out io.Writer, cycles int) {
	const (
		size    = 300
		res     = 0.001
		nframes = 64
		delay   = 8
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randColorIndex()))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func randColorIndex() (index int) {
	return rand.Intn(5)
}
