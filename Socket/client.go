//客户端发送封包
package main

import (
	"./protocol"
	"fmt"
	"net"
	"os"
	"time"
	"bufio"
)

func sender(conn net.Conn) {
	for i := 0; i < 1000; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(protocol.Packet([]byte(words)))
		time.Sleep(1000000000)
	}
	fmt.Println("send over")
}

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	//go sender(conn)
	in := bufio.NewReader(os.Stdin)

	for {
		line, _, _ := in.ReadLine()

		conn.Write(protocol.Packet(line))
		//time.Sleep(1 * 1e9)
	}
}
