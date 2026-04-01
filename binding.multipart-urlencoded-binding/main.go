package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var form loginForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User == "user" && form.Password == "password" {
			c.String(http.StatusOK, "you are logged in")
		} else {
			c.String(http.StatusUnauthorized, "unauthorized")
		}
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
