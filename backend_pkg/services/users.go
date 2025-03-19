package services

import (
	// "fmt"

	"fmt"

	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/Desmond123-arch/PixMorph.git/storage"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	result := storage.Db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUser(username string) (models.User, error) {
	fmt.Println("Finding the user")
	var foundUser models.User
	result := storage.Db.Where("username = ?", username).First(&foundUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	// fmt.Println(foundUser)
	return foundUser, nil
}
