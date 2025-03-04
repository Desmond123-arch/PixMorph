package auth

import "github.com/gin-gonic/gin"

func Create(c *gin.Context) {
	c.JSON(201, gin.H{"Create":"World"})
}

func Login(c *gin.Context) {

	c.JSON(200, gin.H{"Login":"Success"})
}