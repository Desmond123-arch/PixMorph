package services

import (
	// "fmt"
	"fmt"
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/Desmond123-arch/PixMorph.git/storage"
)

func ListAllImages() {
	fmt.Println("Getting all images")
	var images []models.Image
	result := storage.Db.Find(&images)
	if result.Error != nil {
		panic("Error occurring")
	}
	fmt.Println(images)
}
func SearchImage(id string) (models.Image, error){
	var foundImage models.Image
	result := storage.Db.Where("id = ?", id).First(&foundImage)
	if result.Error != nil {
		return models.Image{}, result.Error
	}
	return foundImage, nil
}

// func SaveImage(image *models.Image) error {
// 	r
	
// 	return nil
// }
func DownloadImage(url string) ([]byte, error){
	url = strings.Trim(url, "\n\r")
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var buf bytes.Buffer

	defer response.Body.Close()

	_, err = io.Copy(&buf, response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return buf.Bytes(), nil
}
