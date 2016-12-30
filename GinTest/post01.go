package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"net/http"

)


//// Binding from JSON
//type Login struct {
//	User     string `form:"user" json:"user" binding:"required"`
//	Password string `form:"password" json:"password" binding:"required"`
//}


func main() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(middleware)

	//router.SetHTMLTemplate(tpl)
	router.LoadHTMLGlob("templates/*")
	router.GET("/index",htmlIndex)
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	router.POST("/api/login/",loginhandle)
	router.POST("/api/loginform",loginForm)
	router.Run(":8081")
}

func middleware(c *gin.Context)  {
	//if c.Request.Form == nil {
	//	c.Request.ParseMultipartForm(32 << 20)
	//}
	fmt.Println(c.Request.FormValue("name"))

	fmt.Println("param:,",c.Request.Form, "URL: ",c.Request.URL,"Params:",c.Param)

}

func htmlIndex(c *gin.Context)  {
	c.HTML(http.StatusOK,"temp.tmpl",gin.H{
		"title": "Main website",
	})
}


// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


func loginhandle(c * gin.Context)  {

	//value := c.Request.FormValue("user")
	//fmt.Println(c.Request.FormValue("user"))
	//
	//
	//c.JSON(http.StatusOK, gin.H{"status": value})

	fmt.Println(c.Request.FormValue("user"))
	var json Login
	if c.BindJSON(&json) == nil {
		if json.User == "manu" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}


}

func loginForm(c *gin.Context)  {
	var form Login
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&form) == nil {
		if form.User == "manu" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}



//const tpl = `<!DOCTYPE html>
//<html lang="en">
//<head>
//    <meta charset="UTF-8">
//    <title>Title</title>
//</head>
//<body>
//
//</body>
//</html>`