package newton

import (
	"image/color"
	"math/cmplx"
)

func newtonComplex64(z complex64) color.Color {
	const epsilon = 1e-6
	var roots64 = []complex64{
		1 + 0i,
		-1 + 0i,
		0 + 1i,
		0 - 1i,
	}

	for n := uint8(0); n < iterations; n++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		for i, root := range roots64 {
			if cmplx.Abs(complex128(z-root)) < epsilon {
				return color.RGBA{uint8(255 - contrast*n), uint8((255 - contrast*n) * uint8(i)), 255 - contrast*n, 255}
			}
		}
	}
	return color.Black
}

func newtonComplex128(z complex128) color.Color {
	const epsilon = 1e-6
	for n := uint8(0); n < iterations; n++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		for i, root := range roots {
			if cmplx.Abs(z-root) < epsilon {
				return color.RGBA{uint8(255 - contrast*n), uint8((255 - contrast*n) * uint8(i)), 255 - contrast*n, 255}
			}
		}
	}
	return color.Black
}
