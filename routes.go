package main

import (
	"fmt"
	"net/http"

	goose "github.com/advancedlogic/GoOse"
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/", index)
	router.POST("/article/new", saveArticle)
	router.GET("/article/save", saveHTML)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "Howdy!",
		"payload": articles,
	})
}

func saveArticle(c *gin.Context) {
	url := c.PostForm("url")
	g := goose.New()
	article, err := g.ExtractFromURL(url)
	if err != nil {
		fmt.Println(err)
	}
	tmpArticle := Article{
		Title:       article.Title,
		Content:     article.CleanedText,
		Description: article.MetaDescription,
	}
	articles = append(articles, tmpArticle)
	c.HTML(http.StatusOK, "saved.html", nil)
}

func saveHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "article.save.html", nil)
}
