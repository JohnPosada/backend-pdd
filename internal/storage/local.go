package storage

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LocalStorage struct{}

func (s *LocalStorage) SavePhoto(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["uploads"]

	for _, file := range files {
		log.Println(file.Filename)

		//Upload the file to specific dst.
		c.SaveUploadedFile(file, "./uploads/"+uuid.New().String()+filepath.Ext(file.Filename))
		c.SaveUploadedFile(file, "./"+uuid.New().String()+filepath.Ext(file.Filename))
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
