package transform

import (
	"image"
	"image/color"

	"github.com/Desmond123-arch/pkg/utils"
)

func Linear_transform(img image.Image, new_width int, new_height int) *image.RGBA {
	image_width := img.Bounds().Max.X
	image_height := img.Bounds().Max.Y

	width_scale := float64(image_width-1) / float64(new_width-1)
	height_scale := float64(image_height-1) / float64(new_height-1)
	var alpha float64
	newImage := image.NewRGBA(image.Rect(0, 0, new_width, image_height))
	finalImage := image.NewRGBA(image.Rect(0, 0, new_width, new_height))
	// HORIZONTAL SCALLING
	for y := 0; y < image_height; y++ {
		for xNew := 0; xNew < new_width; xNew++ {
			X0 := int(float64(xNew) * width_scale)
			X1 := utils.Clamp(X0+1, X0+1, image_width-1)
			alpha = (float64(xNew) * width_scale) - float64(X0)
			p1 := make(map[string]float64)
			r, g, b, a := img.At(X0, y).RGBA()
			p1["r"] = float64(r)
			p1["g"] = float64(g)
			p1["b"] = float64(b)
			p1["a"] = float64(a)
			r, g, b, a = img.At(X1, y).RGBA()
			p2 := make(map[string]float64)
			p2["r"] = float64(r)
			p2["g"] = float64(g)
			p2["b"] = float64(b)
			p2["a"] = float64(a)

			rval := uint32(p1["r"] + ((p2["r"] - p1["r"]) * alpha))
			gval := uint32(p1["g"] + ((p2["g"] - p1["g"]) * alpha))
			bval := uint32(p1["b"] + ((p2["b"] - p1["b"]) * alpha))
			aval := uint32(p1["a"] + ((p2["a"] - p1["a"]) * alpha))
			newImage.SetRGBA(xNew, y, color.RGBA{
				R: uint8(rval >> 8),
				G: uint8(gval >> 8),
				B: uint8(bval >> 8),
				A: uint8(aval >> 8),
			})
		}
	}

	//VERTICAL SCALLING
	for x := 0; x < image_width; x++ {
		for y := 0; y < new_height; y++ {
			Y0 := int(float64(y) * height_scale)
			Y1 := utils.Clamp(Y0+1, Y0+1, image_height-1)
			alpha = (float64(y) * height_scale) - float64(Y0)
			p1 := make(map[string]float64)
			r, g, b, a := newImage.At(x, Y0).RGBA()
			p1["r"] = float64(r)
			p1["g"] = float64(g)
			p1["b"] = float64(b)
			p1["a"] = float64(a)

			r, g, b, a = newImage.At(x, Y1).RGBA()
			p2 := make(map[string]float64)
			p2["r"] = float64(r)
			p2["g"] = float64(g)
			p2["b"] = float64(b)
			p2["a"] = float64(a)

			rval := uint32(p1["r"] + ((p2["r"] - p1["r"]) * alpha))
			gval := uint32(p1["g"] + ((p2["g"] - p1["g"]) * alpha))
			bval := uint32(p1["b"] + ((p2["b"] - p1["b"]) * alpha))
			aval := uint32(p1["a"] + ((p2["a"] - p1["a"]) * alpha))
			finalImage.SetRGBA(x, y, color.RGBA{
				R: uint8(rval >> 8),
				G: uint8(gval >> 8),
				B: uint8(bval >> 8),
				A: uint8(aval >> 8),
			})
		}
	}
	return finalImage
}
