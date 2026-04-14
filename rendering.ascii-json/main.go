package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ascii-json", func(c *gin.Context) {
		data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}