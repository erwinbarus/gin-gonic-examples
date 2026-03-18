package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func getWithQueryString(c *gin.Context) {
	fistName := c.Query("first")
	lastName := c.DefaultQuery("last", "Guest")

	c.String(http.StatusOK, "Hello %s %s", fistName, lastName)
}

func main() {
	router := gin.Default();
	router.GET("welcome", getWithQueryString)
	router.Run(":3000")
}