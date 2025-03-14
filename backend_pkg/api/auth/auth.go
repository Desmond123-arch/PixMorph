package auth

import (
	"errors"
	"fmt"
	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/Desmond123-arch/PixMorph.git/services"
	"github.com/Desmond123-arch/PixMorph.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect parameters"})
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		fmt.Printf("%#v\n", err)
		//unwrappedErr := errors.Unwrap(err)
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		} else {
			fmt.Println("Here")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	access_token, err := utils.CreateToken(user.Username)
	refresh_token, err := utils.CreateRefreshToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusCreated, gin.H{"username": user.Username, "access_token": access_token})
	c.SetCookie("refresh_token", refresh_token, 3600*24, "/", "", false, true)
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect parameters"})
		return
	}
	foundUser, err := services.GetUser(user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42P0142P01" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		}
	}
	fmt.Println(foundUser)
	isValid := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	fmt.Println(isValid)
	if isValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}
	access_token, err := utils.CreateToken(user.Username)
	refresh_token, err := utils.CreateRefreshToken(user.Username)
	c.JSON(200, gin.H{"username": foundUser.Username, "access_token": access_token})
	c.SetCookie("refresh_token", refresh_token, 3600*24, "/", "", false, true)

}
