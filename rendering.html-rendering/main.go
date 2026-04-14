package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	tmpl := template.Must(template.ParseGlob("templates/*.tmpl"))
	tmpl = template.Must(tmpl.ParseGlob("templates/*/*.tmpl"))

	router.SetHTMLTemplate(tmpl)
	
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website",
		})
	}) 

	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "Users",
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}