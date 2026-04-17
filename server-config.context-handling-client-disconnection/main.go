package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/api/stream", func(c *gin.Context) {
        ctx := c.Request.Context()

        for i :=0; ;i++ {
            select {
            case <-ctx.Done():
                log.Println("client disconnected, stopping work")
                return
            case <-time.After(1 * time.Second):
                c.SSEvent("message", gin.H{"count": i})
                c.Writer.Flush()
            }
        }
    })

    return router
}

func main() {
    router := setupRouter()
    http.ListenAndServe(":3000", router)
}