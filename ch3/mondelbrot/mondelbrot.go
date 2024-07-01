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
	subsamples             = 4
)

func Mondelbrot(w http.ResponseWriter, r *http.Request) {
	rect := image.Rect(0, 0, width, height)
	rgba := image.NewRGBA(rect)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + xmin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			rgba.Set(px, py, avergeColor(x, y))
		}
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, rgba)
}

func avergeColor(x float64, y float64) color.Color {
	const dx = (xmax - xmin) / width / subsamples
	const dy = (ymax - ymin) / height / subsamples

	var r, g, b, a uint32

	for i := 0; i < subsamples; i++ {
		for j := 0; j < subsamples; j++ {
			sx := x + (float64(i)-0.5)*dx
			sy := y + (float64(i)-0.5)*dy

			c := helpMondelbrot(complex(sx, sy))
			cr, cg, cb, ca := c.RGBA()
			r += cr
			g += cg
			b += cb
			a += ca
		}
	}
	return color.RGBA{
		R: uint8(r / (subsamples * subsamples) >> 8),
		G: uint8(g / (subsamples * subsamples) >> 8),
		B: uint8(b / (subsamples * subsamples) >> 8),
		A: uint8(a / (subsamples * subsamples) >> 8),
	}
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
