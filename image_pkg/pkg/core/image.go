package core

import (
	"image"
)

type ImageProcessor interface {
	Resize(width, height int) error
	Rotate(degrees float64) error

	AdjustBrightness(factor float64) error
	AdjustContrast(factor float64) error
}

type ProcessableImage struct {
	img image.Image
}

type ResizeOptions struct {
	width, height int
}
