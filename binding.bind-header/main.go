package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	var h testHeader
	router.GET("/", func(c *gin.Context) {
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"rate": h.Rate, "domain": h.Domain})
	})
	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
