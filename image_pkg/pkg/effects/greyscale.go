package effects

import (
	"image"
	"image/color"
)

func ApplyGrayColor(img image.Image) *image.RGBA {
	boundsX := img.Bounds().Max.X
	boundsY := img.Bounds().Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, boundsX, boundsY))

	for i := 0; i < boundsX; i++ {
		for j := 0; j < boundsY; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)
			a8 := uint8(a >> 8)
			gray := uint8(0.299*r8 + 0.587*g8 + 0.114*b8)
			newImage.SetRGBA(i, j, color.RGBA{
				R: gray,
				G: gray,
				B: gray,
				A: uint8(a8),
			})
		}
	}
	return newImage
}
