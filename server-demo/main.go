package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/brasbug/log4go"
	"project/server-demo/handlers"
	"project/server-demo/ginpprof"
	"project/server-demo/db"
	"project/server-demo/midware"
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

func init()  {
	//链接mgo
	db.Connect()
}

func main() {

	SetLog()
	defer log.Close()

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(midware.Connect)
	router.Use(midware.Auth())
	router.Use(midware.Middleware)
	router.GET("/get", handlers.GetHandler)
	router.POST("/articls",handlers.Create)
	ginpprof.Wrapper(router)
	router.Run(":8081")


}
