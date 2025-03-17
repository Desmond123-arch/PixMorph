package images

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	gParser "github.com/Desmond123-arch/GParser"
	"github.com/Desmond123-arch/PixMorph.git/models"
	"github.com/Desmond123-arch/PixMorph.git/services"
	"github.com/Desmond123-arch/PixMorph.git/storage"
	"github.com/Desmond123-arch/PixMorph.git/utils"
	"github.com/gin-gonic/gin"
)

type Item struct {
	image os.File
}

func Upload(c *gin.Context) {
	image, err := c.FormFile("image")
	now := time.Now()
	timeString := now.Format("20060102150405")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image invalid"})
	}
	image.Filename = timeString + "-" + filepath.Base(image.Filename)

	val,err := utils.UploadImage(image)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"Image upload failed"})
		return
	}
	url, _ := gParser.Parse(val)
	image_url := fmt.Sprintf("\n%s/object/public/%s\n", os.Getenv("SUPABASE_URL"), url["Key"])
	var upImage models.Image
	rawUsername, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found"})
		return
	}
	username, ok := rawUsername.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username type"})
		return
	}
	user,err  := services.GetUser(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	upImage.UserID = user.ID
	upImage.Url = image_url
	storage.Db.Create(&upImage)
	c.JSON(http.StatusOK, gin.H{"image": "good one"})
}

// func Transform(c *gin.Context) {
// 	image_id := c.Param("id")

// }