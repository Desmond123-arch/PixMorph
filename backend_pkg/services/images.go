package services

import (
	// "fmt"
	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/Desmond123-arch/PixMorph.git/storage"
)


func SearchImage(id string) (models.Image, error){
	var foundImage models.Image
	result := storage.Db.Where("id = ?", id).First(&foundImage)
	if result.Error != nil {
		return models.Image{}, result.Error
	}
	// fmt.Println(foundImage)
	return foundImage, nil
}

// func SaveImage(image *models.Image) error {
// 	r
	
// 	return nil
// }