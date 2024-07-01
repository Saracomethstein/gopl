package newton

import (
	"image"
	"image/png"
	"net/http"
)

func NewtonComplex64Handler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newtonComplex64(z))
		}
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, img)
}

func NewtonComplex128Handler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newtonComplex128(z))
		}
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, img)
}
