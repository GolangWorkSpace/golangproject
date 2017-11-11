
package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Logintime time.Time
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func main() {
	user := User{5, "zhangsan", "pwd", time.Now()}
	data := Struct2Map(user)
	fmt.Println(data)


}