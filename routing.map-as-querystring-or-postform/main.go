package main

import (
	"fmt"
 	"net/http"

  	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v\n", ids, names)
		c.JSON(http.StatusOK, gin.H{
		"ids":   ids,
		"names": names,
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}