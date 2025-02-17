package main

import (
	"fmt"
	"github.com/Desmond123-arch/pkg/transform"
	"github.com/Desmond123-arch/pkg/utils"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	imgFile, err := os.Open("./109918035.jpg")

	if err != nil {
		//FIXME: Change this later
		panic(err)
	}
	defer imgFile.Close()

	img, imageType, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	//newImage := effects.ApplyGrayColor(img)
	newImage := transform.Linear_transform(img, 300, 200)
	imagename := fmt.Sprintf("newImage.%s", imageType)
	newImageFile, err := os.Create(imagename)
	if err != nil {
		panic(err)
	}
	err = utils.WriteToFile(newImage, newImageFile, imageType)
	if err != nil {
		panic(err)
	}
	defer newImageFile.Close()

}
