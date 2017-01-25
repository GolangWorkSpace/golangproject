package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	log "github.com/brasbug/log4go"
	"Project/server-demo/handlers"
	"Project/server-demo/ginpprof"
	"Project/server-demo/db"
	"Project/server-demo/midware"
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

	//链接mgo
	db.Connect()

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(midware.Connect)
	router.Use(midware.ErrorHandler)
	router.Use(midware.Middleware)
	router.GET("/get", handlers.GetHandler)
	router.POST("/articls",handlers.Create)
	ginpprof.Wrapper(router)
	router.Run(":8081")


}
