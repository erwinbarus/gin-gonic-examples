package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "method":"GET" })
}

func posting(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{ "method":"POST" })
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "method":"PUT" })
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "method":"DELETE" })
}

func optioning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "method":"OPTIONS" })
}

func heading(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "method":"HEAD" })
}

func main() {
	router := gin.Default()

	router.GET("/", getting)
	router.POST("/", posting)
	router.PUT("/", putting)
	router.DELETE("/", deleting)
	router.OPTIONS("/", optioning)
	router.HEAD("/", heading)

	router.Run(":3000")
}

