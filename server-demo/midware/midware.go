package midware

import (
	"gopkg.in/gin-gonic/gin.v1"
	log "github.com/brasbug/log4go"
)


func Middleware(c *gin.Context) {
	if c.Request.Form == nil {
		c.Request.ParseMultipartForm(32 << 20)
	}
	log.Info("%s,%s%s",c.Request.Form,c.Request.Host,c.Request.RequestURI)
}
