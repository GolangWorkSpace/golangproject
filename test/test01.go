
package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)

	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:", strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello world")
}


func main() {
	fmt.Println("Hello Golang!")

	http.HandleFunc("/",sayHelloName)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServer: ",err)
	}
}



