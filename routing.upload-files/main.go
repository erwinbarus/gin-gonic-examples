package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
  MaxUploadSize = 1 << 20 // 1 MB
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 1 << 20 // 1 MiB

	router.POST("/upload/single", UploadSingleFile)
	router.POST("/upload/multiple", UploadMultipleFiles)

	return router
}

func UploadSingleFile(c *gin.Context) {
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

    file, err := c.FormFile("file")
    if err != nil {
      if _, ok := err.(*http.MaxBytesError); ok {
        c.JSON(http.StatusRequestEntityTooLarge, gin.H{
          "error": fmt.Sprintf("file too large (max: %d bytes)", MaxUploadSize),
        })
        return
      }
      
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    log.Println(file.Filename)

    dst := filepath.Join("./files/", filepath.Base(file.Filename))
    c.SaveUploadedFile(file, dst)

    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func UploadMultipleFiles(c *gin.Context) {
    form, err := c.MultipartForm()
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    files := form.File["files"]

    for _, file := range files {
      log.Println(file.Filename)

      dst := filepath.Join("./files/", filepath.Base(file.Filename))
      c.SaveUploadedFile(file, dst)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func main() {
	router := SetupRouter()
	router.Run(":3000")
}