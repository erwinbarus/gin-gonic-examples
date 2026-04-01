package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": p.Name, "id": p.ID})
	})
	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
