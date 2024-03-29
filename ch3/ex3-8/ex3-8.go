package main

import (
	"os"
	"image"
	"image/png"

	"github.com/jttait/gopl.io/ch3/fractals"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin

			//z := complex(x, y)
			//img.Set(px, py, fractals.NewtonsComplex128(z))

			//z := complex64(complex(x, y))
			//img.Set(px, py, fractals.NewtonsComplex64(z))

			img.Set(px, py, fractals.NetwonsBigFloat(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}
