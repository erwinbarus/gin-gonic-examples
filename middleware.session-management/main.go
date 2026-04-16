package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    router := gin.Default()

    store := cookie.NewStore([]byte("your-secret-key"))
    router.Use(sessions.Sessions("mysession", store))

    router.GET("/login", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Set("user", "john")
        session.Save()
        c.JSON(http.StatusOK, gin.H{"message": "logged in"})
    })

    router.GET("/profile", func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        if user == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"user": user})
    })

    router.GET("/logout", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Clear()
        session.Save()
        c.JSON(http.StatusOK, gin.H{"message": "logged out"})
    })

    return router
}

func main() {
    router := setupRouter()
    router.Run(":3000")
}