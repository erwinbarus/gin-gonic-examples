package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("healthz", func(c *gin.Context) {
        c.String(http.StatusOK, "ok")
    })

    return router
}

func main() {
    server := &http.Server{
        Addr: ":3000",
        Handler: setupRouter(),
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    server.ListenAndServe()
}