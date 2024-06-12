package main

import (
	"fmt"
	"log"
	"os"
)

// http://www.zzvips.com/article/207448.html
// https://blog.csdn.net/chen1415886044/article/details/104626779
func test()  {
	println("11")
	println("22")
}

func Debug(logName string) {
	// 新建一个日志文件
	logFile,err := os.OpenFile(logName,os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	// 打开文件的错误处理
	if err!= nil {
		fmt.Printf("create file %v error :%v",logName,err)
	}
	// 关闭日志文件
	if logFile != nil {
		defer func(file *os.File) {file.Close()}(logFile)
	}

	debugLog := log.New(logFile,"[DEBUG]",log.Ldate)
	debugLog.SetPrefix("[DEBUG]")
	debugLog.SetFlags(log.Lshortfile)
	//debugLog.Println("this is DEBUG log")

}
func main()  {
	//2023/04/27 19:18:50 日志
	//2023/04/27 19:18:50 日志不换行
	//2023/04/27 19:18:50 日志可格式化

	//log.Println("日志")
	//log.Print("日志不换行")
	//log.Printf("日志可%v","格式化")

	//defer log.Println("结束")
	// 打印完日志后不再输出

	//log.Fatal("日志？")
	//log.Fatalln("日志？？？")
	//log.Fatalf("日志可格式化%v","fatal")

	//log.Panicln("报错信息 。。。")




}
