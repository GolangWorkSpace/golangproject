//客户端发送封包
package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"bufio"
	"./protocol"

)

func main() {

	server := "127.0.0.1:5000"
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
	in := bufio.NewReader(os.Stdin)
	msg := RandString(2040)
	words := protocol.Packet([]byte(msg))
	fmt.Println(len(words), string(words))
	conn.Write(words)



	for  {
		line, _, _ := in.ReadLine()
		fmt.Println(line)
		send(conn,words)
	}




}

func send(conn *net.TCPConn,line []byte)  {
	conn.Write(line)
}



/**
*生成随机字符
**/
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}