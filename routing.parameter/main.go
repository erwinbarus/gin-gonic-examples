package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getWithParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello, %s!", name)
}

func getWithParams(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", getWithParam)
	router.GET("/user/:name/*action", getWithParams)

	router.Run(":3000")
}