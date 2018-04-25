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
	"im_go/libs/proto"

	"im_go/libs/bytes"
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
	//words := protocol.Packet([]byte(msg))
	//fmt.Println(len(words), string(words))
	//conn.Write(words)
	//fmt.Println(utils.Md5Buf(words))
	p := proto.Proto{
		Ver:1,
		Operation:2,
		SeqId:3,
		Body:[]byte(msg),
	}
	var w  bytes.Writer
	p.WriteTo(&w)
	buf := w.Buffer()
	fmt.Println("length",len(buf),string(buf))
	buffer := make([]byte, 1024)

	for  {
		line, _, _ := in.ReadLine()
		fmt.Println(line)
		send(conn,buf)
		n, _ := conn.Read(buffer)
		fmt.Println(string(buffer[:n]))
		buffer = buffer[:0]

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