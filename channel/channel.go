package main

import (
	"fmt"
	"time"
)

const (
	num = 10000000  // 测试1千万次发送和接收
)

func main() {
	TestChan2()
}

func TestChan2() {
	st := time.Now().UnixNano()

	c := make(chan int)

	go func() {
		var n int
		for n = range c {

		}
		fmt.Printf("%d", n)

		fmt.Printf("task TestChan2 cost %d \r\n", (time.Now().UnixNano()-st)/int64(time.Millisecond))

	}()

	for i := 0; i < num; i++ {
		c <- i
	}
	close(c) // 发送完就关闭通道， 接手协程中就会 退出 for range 循环

	// sleep 一段时间，确保接收chan完成
	time.Sleep(3 * time.Second)

}