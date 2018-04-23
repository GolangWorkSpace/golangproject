package main
import (
	"fmt"
	"net/http"
	"time"
	"regexp"
)
func Foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	w.Write([]byte("HelloWorld!"))
	time.Sleep(time.Second * 4)
}
func main() {
	//http.HandleFunc("/hello", Foo)
	//http.ListenAndServe(":8080", nil)

	routerRe    := regexp.MustCompile(`([A-Z]+)[^/]*(/[a-zA-Z/]+)`)
	match := routerRe.FindAllStringSubmatch("GET   /system",-1)
	for _,v := range match[0]{
		fmt.Println(v)
	}

	match = routerRe.FindAllStringSubmatch("POST   /updategroup",-1)
	for _,v := range match[0]{
		fmt.Println(v)
	}


}