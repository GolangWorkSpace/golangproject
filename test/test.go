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


	ip4 := taddr.IP.To4()




}


func httpPost() {
	resp, err := http.Post("http://localhost:3001/",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}