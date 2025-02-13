// implements a blur effect using box blur
package effects

import (
	"github.com/Desmond123-arch/pkg/utils"
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
	Xbound := img.Bounds().Max.X
	Ybound := img.Bounds().Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, Xbound, Ybound))
	ksize := 7
	sigma := 7.0

	kernel := make([][]float64, ksize)
	var kernelSum float64
	for i := range kernel {
		kernel[i] = make([]float64, ksize)
	}
	center := int(math.Floor(float64(len(kernel) / 2)))

	for i := range len(kernel) {
		for j := range len(kernel[i]) {
			x := float64(i - center)
			y := float64(j - center)
			weight := math.Exp(-(x*x + y*y) / (2 * sigma * sigma))
			kernel[i][j] = weight
			kernelSum += weight
		}
	}

	//normalize kernel
	for i := range kernel {
		for j := range kernel[i] {
			kernel[i][j] /= kernelSum
		}
	}

	for x := 0; x < Xbound; x++ {
		for y := 0; y < Ybound; y++ {
			var tempSum RGBASum
			var colorRGBA color.RGBA

			for i := range ksize {
				for j := range ksize {
					srcX := x + i - center
					srcY := y + j - center

					srcX = utils.Clamp(srcX, 0, ksize)
					srcY = utils.Clamp(srcY, 0, ksize)
					r, g, b, a := img.At(srcX, srcY).RGBA()

					tempSum.sumRed += uint32(float64(r) * kernel[i][j])
					tempSum.sumBlue += uint32(float64(b) * kernel[i][j])
					tempSum.sumGreen += uint32(float64(g) * kernel[i][j])
					tempSum.sumAlpha += uint32(float64(a) * kernel[i][j])
				}

			}
			colorRGBA.R = uint8(tempSum.sumRed / 273 >> 8)
			colorRGBA.B = uint8(tempSum.sumBlue / 273 >> 8)
			colorRGBA.G = uint8(tempSum.sumGreen / 273 >> 8)
			colorRGBA.A = uint8(tempSum.sumAlpha / 273 >> 8)
		}
		newImage.SetRGBA(x, y, colorRGBA)
	}
	return newImage
}
