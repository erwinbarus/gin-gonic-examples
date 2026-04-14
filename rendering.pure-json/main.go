package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
 
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"html": "<b>Hello, world!</b>",
		})
	})

	router.GET("/pure-json", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
		"html": "<b>Hello, world!</b>",
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}