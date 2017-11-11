package main

import (
	"github.com/fatih/structs"
	//"net/http"

	"fmt"

)

type Server struct {
	Name        string `json:"name"`
	ID          int
	Enabled     bool
	User       string
}


func main()  {


	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
		User: "a9090909",
	}

	fmt.Println(structs.Map(server))


}
