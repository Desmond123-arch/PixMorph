package main

import (
	"log"

	"github.com/Desmond123-arch/PixMorph.git/middlewares"
	"github.com/joho/godotenv"

	"github.com/Desmond123-arch/PixMorph.git/api/auth"
	"github.com/Desmond123-arch/PixMorph.git/api/imgs"
	"github.com/Desmond123-arch/PixMorph.git/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ENV")
		//FIXME: Handle dotenv not working edge case
		panic("Dotenv is not working correctly")
	}
	
	//DB CONNECTION
	err = storage.Open()
	if err != nil {
		log.Fatal("Database connection failed")
	}

	//AUTH ROUTES
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Create)
		authRoutes.POST("/login", auth.Login)
		authRoutes.GET("/refresh", auth.RefreshToken)
	}
	//IMAGE ROUTES
	imageRoutes := r.Group("/images")
	{
		imageRoutes.POST("/", middlewares.IsAuthorized(),imgs.Upload)
		imageRoutes.POST("/:id/transform", middlewares.IsAuthorized(), imgs.Transform)
	}
	sqlDB, err := storage.Db.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}
	defer sqlDB.Close()
	r.Run()
}
