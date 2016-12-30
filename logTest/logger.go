package main

import (
	"log"
)

func main() {
	Ldefault()
	Ldate()
	Ltime()
	Lmicroseconds()
	Llongfile()
	Lshortfile()
	LUTC()
}

func Ldefault() {
	log.Println("这是默认的格式\n")
	//log.trace("adss")

}

func Ldate() {
	log.SetFlags(log.Ldate)
	log.Println("这是输出日期格式\n")
}

func Ltime() {
	log.SetFlags(log.Ltime)
	log.Println("这是输出时间格式\n")
}

func Lmicroseconds() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("这是输出微秒格式\n")
}

func Llongfile() {
	log.SetFlags(log.Llongfile)
	log.Println("这是输出路径+文件名+行号格式\n")
}

func Lshortfile() {
	log.SetFlags(log.Lshortfile)
	log.Println("这是输出文件名+行号格式\n")
}

func LUTC() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile)
	log.Println("这是输出 使用标准的UTC时间格式 格式\n")
}

//fatal:非常严重的错误，导致系统终止。期望这类信息被立即显示到状态控制台上。
//Error:其他运行期错误或不是预期的条件。期望这类信息被立即显示到状态控制台上。
//Warn:使用了不赞成使用的API、非常拙劣的使用了API、‘几乎就是错误’其他运行期不合需要和不合预期的状态（但没必要将其称为错误）。期望这类信息被立即显示到状态控制台上。
//Info:运行期产生的有意义的事件。期望这类信息被立即显示到状态控制台上。
//Debug:系统流程中的细节信息。期望这类信息仅被写入日志文件中。
//Trace:更加细节的信息。期望这类信息仅被写入日志文件中。