package images

import (
	"fmt"
	"github.com/Desmond123-arch/PixMorph.git/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
	fmt.Println(image.Filename)

	url,err := utils.UploadImage(image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Image upload failed"})
	}
	fmt.Print("The key is ")
	fmt.Println(url.Key)
	//savePath := "./uploads/" + filename
	//
	//if err := c.SaveUploadedFile(image, savePath); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"image": "good"})
}
