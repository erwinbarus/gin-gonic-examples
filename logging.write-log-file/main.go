package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() (*gin.Engine, *os.File) {
    gin.DisableConsoleColor()

    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    return router, f
}

func main() {
    router, _ := setupRouter()
    router.Run(":3000")
}