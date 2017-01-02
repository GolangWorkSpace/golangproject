package handlers

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	log "github.com/brasbug/log4go"

)

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	log.Info(value)
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}