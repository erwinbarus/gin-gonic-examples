package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("public") 
	{
		public.GET("health", healthChecker)
	}

	private := router.Group("private")
	private.Use(authRequired()) 
	{
		private.GET("profile", getProfile)
		private.GET("setting", getSetting)
	}

	return router
}

func healthChecker(c *gin.Context) {
	c.String(http.StatusOK, "Healthy")
}

func getProfile(c *gin.Context) {
	c.String(http.StatusOK, "View profile here!")
}

func getSetting(c *gin.Context) {
	c.String(http.StatusOK, "View setting here!")
}

func authRequired() gin.HandlerFunc  {
	return func(c *gin.Context) {
		// check session here
		c.Next()
	}
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}