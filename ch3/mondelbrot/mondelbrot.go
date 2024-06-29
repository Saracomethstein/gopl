package mondelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func Mondelbrot(w http.ResponseWriter, r *http.Request) {
	rect := image.Rect(0, 0, width, height)
	rgba := image.NewRGBA(rect)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + xmin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			rgba.Set(px, py, helpMondelbrot(z))
		}
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, rgba)
}

func helpMondelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return color.RGBA{150 - 10*n, 40 - 5*n, 120 + 2*n, 255}
		}
	}
	return color.Black
}
