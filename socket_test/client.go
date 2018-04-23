package main

import (
	"net"
	"fmt"
)

func StartClient1() {
	tcpAddress, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1300")
	if err != nil {
		//errs.Error_exit(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		//errs.Error_exit(err)
	}

	writeChan := make(chan []byte, 1024)
	readChan := make(chan []byte, 1024)

	go writeConnection(conn, writeChan)
	go readConnection(conn, readChan)

	//go handleReadChannel(readChan)

	for {
		var s string
		fmt.Scan(&s)
		writeChan <- []byte(s)
	}

}

func readConnection(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			//errs.Error_print(err)
			return
		}
		println("Received from:", conn.RemoteAddr(), string(buffer[:n]))
		//channel <- buffer[:n]
	}

}

func writeConnection(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()
	for {
		select {
		case data := <- channel:
			_, err := conn.Write(data)
			if err != nil {
				//errs.Error_exit(err)
			}
			println("Write to:", conn.RemoteAddr(), string(data))
		}
	}
}

func main()  {
	StartClient1()
}