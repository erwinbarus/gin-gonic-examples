package main

import (
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err

            c.JSON(http.StatusInternalServerError, gin.H{
                "success": false,
                "message": err.Error(),
            })
        }
    }
}

func setupRouter() *gin.Engine {
    router := gin.New()
    router.Use(ErrorHandler())

    router.GET("/ok", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "Everything is fine!",
        })
    })

    router.Get("/error", func(c *gin.Context) {
        c.Error(errors.New("something went wrong"))
    })

    return router
}

func main() {
    router := setupRouter()
    router.Run(":3000")
}