package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"Project/server-demo/models"
	"gopkg.in/mgo.v2"
)

// New article
func New(c *gin.Context) {
	article := models.Article{}
	c.HTML(http.StatusOK, "articles/form", gin.H{
		"title":   "New article",
		"article": article,
	})
}
// Create an article
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}
	err = db.C(models.CollectionArticle).Insert(article)
	if err != nil {
		c.Error(err)
	}
	//c.Redirect(http.StatusMovedPermanently, "/articles")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": "插入成功",
	})
}
