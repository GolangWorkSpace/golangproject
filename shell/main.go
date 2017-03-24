package main

import (
	"os/exec"
	"fmt"
)


func main() {
	////执行【ls /】并输出返回文本
	//f, err := exec.Command("ls", "/Library").Output()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//执行【ls /】并输出返回文本
	f, err := exec.Command("sudo", "mkdir /Library/abc").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(f))
}
