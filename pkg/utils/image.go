package utils

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func WriteToFile(img *image.RGBA, w io.Writer, imageType string) error {
	if imageType == "png" {
		return png.Encode(w, img)
	} else if imageType == "jpeg" {
		return jpeg.Encode(w, img, nil)
	}
	return errors.New("unsupported image type")
}
