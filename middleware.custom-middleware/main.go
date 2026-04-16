package main

import (
    "log"
    "time"
    "net/http"

    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        c.Set("example", 12345)

        c.Next()

        latency := time.Since(t)
        log.Print(latency)

        status := c.Writer.Status()
        log.Println(status)
    }
}

func setupRouter() *gin.Engine {
    router := gin.New()
    router.Use(Logger())

    router.GET("/test", func(c *gin.Context) {
        example := c.MustGet("example")
        c.JSON(http.StatusOK, example)
    })

    return router
}

func main() {
    router := setupRouter()
    router.Run(":3000")
}