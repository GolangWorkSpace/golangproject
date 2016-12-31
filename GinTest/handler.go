package  GinTest

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
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
