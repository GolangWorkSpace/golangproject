package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"net/http"
	"os"
	"io"
	"log"

)

func main() {


	//初始化命令行参数
	//flag.Parse()
	//
	//glog.Info("hello, glog")
	//glog.Warning("warning glog")
	//glog.Error("error glog")
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(Middleware)
	//router.Use(middle001ware)

	router.POST("/form_post", post1handler)
	router.POST("/posthandle",PostHandler)
	router.POST("/upload",uploadHandler)
	router.GET("/get",GetHandler)
	//router.POST("post",middle001ware)
	router.Run(":8081")





}

func Middleware(c *gin.Context) {
	if c.Request.Form == nil {
		c.Request.ParseMultipartForm(32 << 20)
	}
	fmt.Println("param=",c.Request.Form, "path=",c.Request.URL)



}


func uploadHandler(c *gin.Context)  {
	file, header , err := c.Request.FormFile("upload")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/"+filename+".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	//c.JSON(http.StatusProcessing,gin.H{{
	//	"message":filename,
	//}})
}

func post1handler(c *gin.Context)  {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}


func PostHandler(c *gin.Context) {
	value := c.PostForm("name")
	//fmt.Println(value)
	//type JsonHolder struct {
	//	Id   int    `json:"id"`
	//	Name string `json:"name"`
	//}
	//holder := JsonHolder{Id: 1, Name: value}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, value)
	return
}


//
//
//func main() {
//
//	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
//	router := gin.Default()    //获得路由实例
//
//	//添加中间件
//	//router.Use(Middleware)
//	//注册接口
//	router.GET("/simple/server/get", GetHandler)
//	router.POST("/simple/server/post", PostHandler)
//	router.PUT("/simple/server/put", PutHandler)
//	router.DELETE("/simple/server/delete", DeleteHandler)
//	//监听端口
//	//http.ListenAndServe(":8080", router)
//	router.Run(":8080")
//
//}
//

//
//func GetHandler(c *gin.Context) {
//	value, exist := c.GetQuery("key")
//	if !exist {
//		value = "the key is not exist!"
//	}
//	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
//	return
//}
//func PostHandler(c *gin.Context) {
//	value := c.Param("name")
//
//	type JsonHolder struct {
//		Id   int    `json:"id"`
//		Name string `json:"name"`
//	}
//	holder := JsonHolder{Id: 1, Name: value}
//	//若返回json数据，可以直接使用gin封装好的JSON方法
//	c.JSON(http.StatusOK, holder)
//	return
//}
//func PutHandler(c *gin.Context) {
//	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
//	return
//}
//func DeleteHandler(c *gin.Context) {
//	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
//	return
//}
