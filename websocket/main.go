package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start app")
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":4001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg + "_server")
		if err = websocket.Message.Send(ws, msg+"_server"); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
