package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/bind", func(c *gin.Context) {
		objA := formA{}

		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched formA", "foo": objA.Foo})
			return
		}

		objB := formB{}

		if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched formB", "bar": objB.Bar})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "request body did not match any known format"})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
