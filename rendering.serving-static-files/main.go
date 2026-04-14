package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFS("/public", http.Dir("./public"))
	router.StaticFile("/robots.txt", "robots.txt")

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}