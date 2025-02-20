package transform

import (
	"image"
	"image/color"
	"math"
)

func RotateImage(img image.Image, angle float64) *image.RGBA {
	angle = angle * math.Pi / 180
	origin_width := img.Bounds().Dx()
	origin_height := img.Bounds().Dy()

	corners := [4]struct{ x, y float64 }{
		{0, 0}, {float64(origin_width), 0},
		{0, float64(origin_height)}, {float64(origin_width), float64(origin_height)},
	}
	var minX, minY, maxX, maxY float64
	minX, minY = math.MaxFloat64, math.MaxFloat64
	maxX, maxY = -math.MaxFloat64, -math.MaxFloat64

	for _, c := range corners {
		rotX := c.x*math.Cos(angle) - c.y*math.Sin(angle)
		rotY := c.x*math.Sin(angle) + c.y*math.Cos(angle)

		if rotX < minX {
			minX = rotX
		}
		if rotY < minY {
			minY = rotY
		}
		if rotX > maxX {
			maxX = rotX
		}
		if rotY > maxY {
			maxY = rotY
		}
	}
	newWidth := int(math.Ceil(maxX - minX))
	newHeight := int(math.Ceil(maxY - minY))
	newImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	centerX, centerY := float64(origin_width)/2, float64(origin_height)/2
	newCenterX, newCenterY := float64(newWidth)/2, float64(newHeight)/2

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			origX := (float64(x)-newCenterX)*math.Cos(-angle) - (float64(y)-newCenterY)*math.Sin(-angle) + centerX
			origY := (float64(x)-newCenterX)*math.Sin(-angle) + (float64(y)-newCenterY)*math.Cos(-angle) + centerY
			if origX >= 0 && origX < float64(origin_width) && origY >= 0 && origY < float64(origin_height) {
				r, g, b, a := img.At(int(origX), int(origY)).RGBA()
				newImage.SetRGBA(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
			}
		}
	}
	return newImage
}
