package main

import (
	"net/http"
	"github.com/gin-gonic/gin"	
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("form-post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}