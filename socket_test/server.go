package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"project/socket_test/protocol"
)

func main() {
	netListen, err := net.Listen("tcp", ":5000")
	CheckError(err)

	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}



func handleConnection(conn net.Conn) {
	allbuf := make([]byte, 0)
	buffer := make([]byte, 1024)
	length := 0
	headerlen := 0
	type msg chan []byte

	for {
		readLen, err := conn.Read(buffer)
		if err == io.EOF {
			//conn 连接中断
			fmt.Println("err",err)
			break
		}
		if length == 0 && headerlen == 0 && len(allbuf) == 0{
			headers ,err := protocol.UnPacketHeader(buffer)
			if err == nil {
				fmt.Println("a new in msg")
				length,_ = strconv.Atoi(string(headers[2]))
				headerlen =  len(headers[0])
				allbuf = buffer[headerlen:]
			}
		}else {
			allbuf = append(allbuf,buffer[:readLen]...)
		}
		fmt.Println("readLen: ", readLen, len(allbuf),"head:",string(allbuf[0:24]),"bufferLen:",len(buffer))
		if len(allbuf) == length {
			length = 0
			allbuf = allbuf[:0]
			buffer = make([]byte, 1024)
		}
	}
}



func bufferReadEnd(buffer []byte)bool  {


	return false
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}