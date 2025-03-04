package storage

import (
	"log"
	"os"

	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var Db *gorm.DB

func Open() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ENV")
		return err
	}
	connUrl := os.Getenv("DATABASE_URL")

	Db, err = gorm.Open(postgres.Open(connUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("An Error occured")
		return err
	}
	Db.AutoMigrate(&models.User{}, &models.Image{})
	return nil
}