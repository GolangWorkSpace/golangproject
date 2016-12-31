package handler

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func post001handler(c *gin.Context)  {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
