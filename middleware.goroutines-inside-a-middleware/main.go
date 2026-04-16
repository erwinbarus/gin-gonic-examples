package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    router := gin.New()
    router.Use(gin.Logger())

    router.GET("/long_async", func(c *gin.Context) {
        // create a copy of read-only context to prevent racing condition
        cCp := c.Copy()

        go func() {
            time.Sleep(5 * time.Second)
            log.Println("Done! in path " + cCp.Request.URL.Path)
        }()

        c.String(http.StatusOK, "Done!")
    })

    router.GET("/long_sync", func(c *gin.Context) {        
        time.Sleep(5 * time.Second)

        log.Println("Done! in path " + c.Request.URL.Path)

        c.String(http.StatusOK, "Done!")
    })

    return router
}

func main() {
    router := setupRouter()
    router.Run(":3000")
}