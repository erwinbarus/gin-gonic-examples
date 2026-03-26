package main

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name		string 		`form:"name"`
	Address		string 		`form:"address"`
	Birthday 	time.Time 	`form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("testing", startPage)
	router.POST("testing", startPage)
	return router
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name": person.Name,
		"address": person.Address,
		"birthday": person.Birthday,	
	})
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}