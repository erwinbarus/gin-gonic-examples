package main

import (
	"net/http"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"
)

func SecurityHeaders() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
        c.Header("Referrer-Policy", "strict-origin")
        c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
        c.Next()
    }
}

func setupRouter() *gin.Engine {
    router := gin.New()

    router.Use(gin.Logger())

    router.Use(ginhelmet.Default())
    router.Use(SecurityHeaders())

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, "pong")
    })
    return router
}

func main() {
    router := setupRouter()
    router.Run(":3000")
}