package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"hey"})
	})

	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message":"hey"})
	})

	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message":"hey"})
	})

	router.GET("/toml", func(c *gin.Context) {
		c.TOML(http.StatusOK, gin.H{"message":"hey"})
	})

	router.GET("/proto", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		
		data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
		}
		
		c.ProtoBuf(http.StatusOK, data)
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}