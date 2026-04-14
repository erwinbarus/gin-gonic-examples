package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/secure-json", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names)
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}