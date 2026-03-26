package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Login struct {
	User 		string	`form:"user" json:"user" xml="user" binding:"required"`
	Password 	string	`form:"password" json:"password" xml="password" binding:"required"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/loginJSON", jsonLogin)

	return router
}

func jsonLogin(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "missing field"})
		return
	}

	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}