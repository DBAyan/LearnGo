package main

import "fmt"

func main(){
	//print 直接输出内容，不能格式化输出，也不能换行
	fmt.Print("不会换行！")
	// 相比Print ，Println会换行 ，相当于末尾加\n
	fmt.Println("会换行")
	name := "yhh"
	// Printf 支持格式化输出字符串
	fmt.Printf("hello %s",name)
	// Fprintf

	// Sprintf

}
