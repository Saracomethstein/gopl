package newton

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
	iterations             = 200
	contrast               = 15
)

var roots = []complex128{
	1 + 0i,
	-1 + 0i,
	0 + 1i,
	0 - 1i,
}

func NewtonMondelbrot(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, img)
}

func newton(z complex128) color.Color {
	const epsilon = 1e-6
	for n := uint8(0); n < iterations; n++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		for i, root := range roots {
			if cmplx.Abs(z-root) < epsilon {
				return color.RGBA{uint8(255 - contrast*n),
					uint8((255 - contrast*n) * uint8(i)),
					255 - contrast*n, 255}
			}
		}
	}
	return color.Black
}
