package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)




func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}