package handlers

import (
	"fmt"
	"net/http"
	log "github.com/brasbug/log4go"

	"github.com/gin-gonic/gin"
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