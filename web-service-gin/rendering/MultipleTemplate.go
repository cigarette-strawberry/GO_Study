package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/index.html")
	r.AddFromFiles("article", "templates/base.html")
	return r
}

func MultipleTemplate(c *gin.Context) {
	c.HTML(200, "index", gin.H{"title": "Home"})
	c.HTML(200, "article", gin.H{"title": "Article"})
}
