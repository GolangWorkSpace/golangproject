package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)


type User struct{
	gorm.Model
	Name string
	Age int
	Birthday int64
}

func main()  {

	db, err := gorm.Open("mysql", "test:test@tcp(127.0.0.1:3306)/orm?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now().Unix()}
	db.CreateTable(&User{})

	fmt.Println(db.NewRecord(user))

	fmt.Println(db.Create(&user))

	fmt.Println(db.NewRecord(user)) // => 创建`user`后返回`false
}