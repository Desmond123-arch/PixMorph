package main

import (
	"log"

	"github.com/Desmond123-arch/PixMorph.git/api/auth"
	"github.com/Desmond123-arch/PixMorph.git/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//DB CONNECTION
	err := storage.Open()
	if err != nil {
		log.Fatal("Database connection failed")
	}

	//AUTH ROUTES
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Create)
		authRoutes.POST("/login", auth.Login)
	}
	sqlDB, err := storage.Db.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}
	defer sqlDB.Close()
	r.Run()
}
