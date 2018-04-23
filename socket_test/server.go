package main

import (
	"net"
	"fmt"
	"log"
)

var client_num int = 0

func StartServer1() {
	l, err := net.Listen("tcp", ":1300")
	if err != nil {
		//errs.Error_exit(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		//conn.SetReadDeadline()
		if err != nil {
			//errs.Error_print(err)
			continue
		}
		client_num++
		fmt.Printf("A new Connection %d.\n", client_num)
		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	defer closeConnection(conn)

	readChannel := make(chan []byte, 1024)
	writeChannel := make(chan []byte, 1024)

	go readConnection(conn, readChannel)
	go writeConnection(conn, writeChannel)


	for {
		select {
		case data := <- readChannel:
			if string(data) == "bye" {
				return
			}
			writeChannel <- append([]byte("Back:"), data...)
		}
	}

}

func writeConnection(conn net.Conn, channel chan []byte) {
	for {
		select {
		case data := <- channel:
			println("Write:", conn.RemoteAddr().String(), string(data))
			_, err := conn.Write(data)
			if err != nil {
				//errs.Error_print(err)
				return
			}
		}
	}

}
func readConnection(conn net.Conn, channel chan []byte) {

	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			//errs.Error_print(err)
			channel <- []byte("bye")	//这里须要进一步改进！
			break
		}
		println("Recei:", conn.RemoteAddr().String(), string(buffer[:n]))
		channel <- buffer[:n]
	}
}

func closeConnection(conn net.Conn) {

	conn.Close()
	client_num--
	fmt.Printf("Now, %d connections is alve.\n", client_num)
}





func StartTCPServer()(err error)  {
	var (
		bind     string
		listener *net.TCPListener
		addr     *net.TCPAddr
	)
	bind     =  "127.0.0.1:1300"
	if addr, err = net.ResolveTCPAddr("tcp",bind) ;err !=nil{
		log.Printf("net.ResolveTCPAddr(\"tcp4\", \"%s\") error(%v)", bind, err)
		return
	}
	if listener,err = net.ListenTCP("tcp",addr); err != nil {
		log.Printf("net.ListenTCP(\"tcp4\", \"%s\") error(%v)", bind, err)
		return
	}
	return
}

func acceptTCP(lis *net.TCPListener)  {
	var (
		conn *net.TCPConn
		err  error
	)
	for {
		if conn, err = lis.AcceptTCP(); err != nil {
			// if listener close then return
			log.Printf("listener.Accept(\"%s\") error(%v)", lis.Addr().String(), err)
			return
		}
		if err = conn.SetKeepAlive(false); err != nil {
			log.Printf("conn.SetKeepAlive() error(%v)", err)
			return
		}
		if err = conn.SetReadBuffer(1024); err != nil {
			log.Printf("conn.SetReadBuffer() error(%v)", err)
			return
		}
		if err = conn.SetWriteBuffer(1024); err != nil {
			log.Printf("conn.SetWriteBuffer() error(%v)", err)
			return
		}

		go handlerConnection(conn)
	}

}



func main()  {
	StartServer1()
}