package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	
	router.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})

	var fs http.FileSystem = http.Dir("./")
	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("fs/file.go", fs)
	})

	router.GET("/download", func(c *gin.Context) {
		c.FileAttachment("./local/report-2024-q1.xlsx", "quarterly-report.xlsx")
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}