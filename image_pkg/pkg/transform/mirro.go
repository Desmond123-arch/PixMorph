package transform

import (
	"fmt"
	"image"
	"image/color"
)

func Mirro_image(img image.Image) *image.RGBA {
	image_width := img.Bounds().Max.X
	image_height := img.Bounds().Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, image_width, image_height))
	for y := 0; y < image_height; y++ {
		for x := 0; x < image_width/2; x++ {

			mirroredX := image_width - x - 1
			fr, fg, fb, fa := img.At(x, y).RGBA()
			sr, sg, sb, sa := img.At(mirroredX, y).RGBA()
			newImage.SetRGBA(mirroredX, y, color.RGBA{
				R: uint8(fr >> 8), G: uint8(fg >> 8), B: uint8(fb >> 8), A: uint8(fa >> 8),
			})
			newImage.SetRGBA(x, y, color.RGBA{
				R: uint8(sr >> 8), G: uint8(sg >> 8), B: uint8(sb >> 8), A: uint8(sa >> 8),
			})

		}
	}
	fmt.Println(image_width, image_height)

	return newImage
}
