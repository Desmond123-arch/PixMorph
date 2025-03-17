package auth

import (
	"errors"
	"fmt"
	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/Desmond123-arch/PixMorph.git/services"
	"github.com/Desmond123-arch/PixMorph.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		//unwrappedErr := errors.Unwrap(err)
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		} else {
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
	foundUser, err := services.GetUser(user.Username)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42P0142P01" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		}
	}
	isValid := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	
	if isValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}
	access_token, err := utils.CreateToken(user.Username)
	refresh_token, err := utils.CreateRefreshToken(user.Username)
	c.SetCookie("refresh_token", refresh_token, 3600*24, "/", "", false, true)
	c.JSON(200, gin.H{"username": foundUser.Username, "access_token": access_token})
}

func RefreshToken(c *gin.Context) {
	refresh_token, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
	}
	claim, err := utils.VerifyToken(refresh_token)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid refresh token"})
	}
	if claims, ok := claim.Claims.(jwt.MapClaims); ok && claim.Valid {
		username, exists := claims["username"].(string)
		if !exists {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid refresh token"})
		} else {
			access_token, err := utils.CreateToken(username)
			new_refresh_token, err := utils.CreateRefreshToken(username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			c.SetCookie("refresh_token", new_refresh_token, 3600*24, "/", "", false, true)
			c.JSON(http.StatusOK, gin.H{"access_token": access_token})
		}
	}
}
