package surfacepac

import (
	"fmt"
	"math"
	"net/http"
	"sync"
)

var Cells int = 100

const (
	width, height = 2500, 1400
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	standart      = 1
	box           = 2
	cup           = 3
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var mu sync.Mutex

func Surface(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; background-color: black; stroke-width: 0.3' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < int(Cells); i++ {
		for j := 0; j < int(Cells); j++ {
			ax, ay, az, ok1 := corner(i+1, j)
			bx, by, bz, ok2 := corner(i, j)
			cx, cy, cz, ok3 := corner(i, j+1)
			dx, dy, dz, ok4 := corner(i+1, j+1)

			if ok1 && ok2 && ok3 && ok4 {
				color := colorSVG((az + bz + cz + dz) / 4)
				fmt.Fprintf(w, "<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/float64(Cells) - 0.5)
	y := xyrange * (float64(j)/float64(Cells) - 0.5)

	z := f(x, y, 1) // basic figure

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy) {
		return 0, 0, 0, false
	}

	return sx, sy, z, true
}

func f(x, y float64, figure int) float64 {
	var formula float64
	switch figure {
	case 1:
		r := math.Hypot(x, y)
		formula = math.Sin(r) / r
		break
	case 2:
		formula = math.Sin(x) * math.Cos(y)
		break
	case 3:
		sigma := 10.0
		formula = math.Exp(-(x*x + y*y) / (2 * sigma * sigma))
		break
	}
	return formula
}

func colorSVG(z float64) string {
	minZ, maxZ := -1.0, 1.0
	colorZ := (z - minZ) / (maxZ - minZ)

	red := uint8(255 * colorZ)
	blue := uint8(255 * (1 - colorZ))

	return fmt.Sprintf("#%02x00%02x", red, blue)
}
