package main

import (
	"log"

	"github.com/Desmond123-arch/PixMorph.git/api/auth"
	"github.com/Desmond123-arch/PixMorph.git/storage"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello":"world"})
	})
	//DB CONNECTION
	err := storage.Open()
	if (err != nil) {
		log.Fatal("Database connection failed");
	}
	//AUTH ROUTES
	authRoutes := r.Group("/auth")
	{
		authRoutes.GET("/register", auth.Create)
		authRoutes.POST("/login", auth.Login)
	}
	sqlDB, err := storage.Db.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}
	defer sqlDB.Close()
	r.Run()
}