package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	router.POST("/", func(c *gin.Context) {
		var fakeForm myForm
		if err := c.ShouldBind(&fakeForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
