package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	log "github.com/brasbug/log4go"
	"Project/server-demo/handlers"
	"Project/server-demo/ginpprof"



)

func SetLog() {
	w := log.NewFileWriter()
	w.SetPathPattern("./log/log-%Y%M%D.log")
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	log.SetLevel(log.DEBUG)
	log.SetLayout("2006-01-02 15:04:05")
}

func main() {

	SetLog()
	defer log.Close()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(Middleware)
	router.GET("/get", handlers.GetHandler)
	ginpprof.Wrapper(router)
	router.Run(":8081")
}

func Middleware(c *gin.Context) {
	if c.Request.Form == nil {
		c.Request.ParseMultipartForm(32 << 20)
	}
	log.Info("%s,%s%s",c.Request.Form,c.Request.Host,c.Request.RequestURI)
}