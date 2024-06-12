package main

import "time"

/*
区分并发与并行

并发 处理多个任务的能力
并行 是指同时处理多个任务

并发 代码结构
并行 关乎执行

C语言的并发
多线程模型

缺点 ：
复杂，难于规模化

Go语言如何实现并发
goroutine 一由 Go 运行时（runtime）负责调度的、轻量的用户级线程

优势
占用资源少
并发控制在用户层 而不是在操作系统层
在语言层面 而不是通过标准库提供
语言内置 channel 作为 goroutine 间通信原语，为并发设计提供了强大支撑。


无论是 Go 自身运行时代码还是用户层 Go 代码，都无一例外地运行在 goroutine 中。

创建 goroutine

go func

退出 goroutine
goroutine 的执行函数的返回，就意味着 goroutine 退出。

goroutine的通信


*/

func hello()  {
	println("Hello Goroutine!")
}

func main()  {
	println("call main")

	go hello()
	time.Sleep(100*time.Millisecond)
	println("end main")
}
