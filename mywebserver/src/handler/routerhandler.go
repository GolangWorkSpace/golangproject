package handler


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context) {
	name := c.PostForm("username")
	password   := c.PostForm("password")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"name": name,
		"password":    password,
	})
}