
package main

import (
	"fmt"
	"net"

	"strings"
	"strconv"
	"bytes"
)


func main() {
	address := fmt.Sprintf("2.0.0.0:%d",9000)
	tcpAddr,_ := net.ResolveTCPAddr("tcp",address)
	//address,_ := net.ResolveIPAddr("tcp","0.0.0.0:8000")
	ip4 := tcpAddr.IP.To4()
	public_ip := int32(ip4[0]) << 24 | int32(ip4[1]) << 16 | int32(ip4[2]) << 8 | int32(ip4[3])
	fmt.Println(public_ip)

	fmt.Println(StringIpToInt("2.4.0.0"))
}



func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

