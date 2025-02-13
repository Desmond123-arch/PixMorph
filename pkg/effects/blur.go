// implements a blur effect using box blur
package effects

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

type RGBASum struct {
	sumRed   uint32
	sumBlue  uint32
	sumGreen uint32
	sumAlpha uint32
}

func ApplyBoxBlur(img image.Image) *image.RGBA {
	possibleMoves := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	x := img.Bounds().Max.X
	y := img.Bounds().Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, x, y))
	for i := range y {
		for j := range x {
			r, g, b, a := img.At(i, j).RGBA()
			var tempSum RGBASum
			var colorRGBA color.RGBA

			if !(i == 0 || i == 399) && !(j == 0 || j == 399) {
				for k := range possibleMoves {
					dx := i + possibleMoves[k][0]
					dy := j + possibleMoves[k][1]
					r, g, b, a := img.At(dx, dy).RGBA()
					tempSum.sumRed += r
					tempSum.sumBlue += b
					tempSum.sumGreen += g
					tempSum.sumAlpha += a
				}
				colorRGBA.R = uint8(tempSum.sumRed / 9 >> 8)
				colorRGBA.B = uint8(tempSum.sumBlue / 9 >> 8)
				colorRGBA.G = uint8(tempSum.sumGreen / 9 >> 8)
				colorRGBA.A = uint8(tempSum.sumAlpha / 9 >> 8)
			} else {
				colorRGBA.R = uint8(r >> 8)
				colorRGBA.B = uint8(b >> 8)
				colorRGBA.G = uint8(g >> 8)
				colorRGBA.A = uint8(a >> 8)
			}
			newImage.SetRGBA(i, j, colorRGBA)
		}
	}
	return newImage
}

func ApplyGaussianBlur(img image.Image) *image.RGBA {
	ks := 15
	k := make([]float64, ks*ks)
	for i := 0; i < ks; i++ {
		for j := 0; j < ks; j++ {
			k[i*ks+j] = math.Exp(-(math.Pow(float64(i)-15/2, 2)+math.Pow(float64(j)-15/2, 2))/(2*math.Pow(15/2, 2))) / 2

		}
	}
	dst := image.NewRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			var r, g, b, a float64
			for ky := 0; ky < ks; ky++ {
				for kx := 0; kx < ks; kx++ {
					c := img.At(x+kx-ks/2, y+ky-ks/2)
					r1, g1, b1, a1 := c.RGBA()

					k := k[ky*ks+kx]

					r += float64(r1) * k
					g += float64(g1) * k
					b += float64(b1) * k
					a += float64(a1) * k

				}
			}
			dst.SetRGBA(x, y, color.RGBA{R: uint8(r / 273), G: uint8(g / 273), B: uint8(b / 273), A: uint8(a / 273)})
		}
	}
	return dst
}
