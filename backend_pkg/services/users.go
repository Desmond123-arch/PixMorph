package services

import (
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
func GetUser(user models.User) (models.User, error) {
	var foundUser models.User
	result := storage.Db.Where("username = ?", user.Username).First(&foundUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	fmt.Println(foundUser)
	return foundUser, nil
}
