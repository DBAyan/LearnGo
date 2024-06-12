package main

import (
	"errors"
	"fmt"
)

// 一个目录下只能有一个main 函数 ，想要写多个main 函数
func defaultValues()  {
	var s string
	var i int
	var b bool
	var f float32
	var c complex64
	var arr []int
	var r rune
	var B byte
	fmt.Printf("string default value is *%s*\n", s) // 空字符串
	fmt.Printf("int default value is %d\n", i) // 0
	fmt.Printf("bool default value is %t\n", b) //false
	fmt.Printf("float default value is %f，%e, %g\n", f,f,f) // 0.000000
	fmt.Printf("conplex64 default value is %v\n", c) //
	fmt.Printf("array default value is %v,arr is nil %t\n", arr, arr == nil) // []
	fmt.Printf("rune default value is %d\n", r) // 0
	fmt.Printf("byte default value is %d\n", B) // 0
	/*
		string default value is
		int default value is 0
		bool default value is falsfloat default value is 0.000000
		conplex64 default value is (0+0i)
		array default value is []
		rune default value is 0
		byte default value is 0
	*/
}


func scope()  {
	// 类型的上限
	var b byte = 255 // byte类型上线255
	var a int8 = 127 // int8上限127

	var o bool = true // 声明并初始化
	var m bool // 只声明不赋值，默认零值
	n := false // 通过类型推断 ，短变量声明
	fmt.Println(a, b, o, m,n)

	// 复数
	
	// 指针

	// error类型
	var e error
	e = errors.New("错误信息")
	fmt.Printf("%v\n",e)
	fmt.Printf("%+v\n",e)
	fmt.Printf("%#v\n",e)

	//  %v %+v %#v的区别
	type User struct {
		name string
		age int
	}
	zs := User{name: "张三", age: 30}
	fmt.Printf("%v\n",zs)
	fmt.Printf("%+v\n",zs)
	fmt.Printf("%#v\n",zs)
	// {张三 30}  %v 只输出了结构体的值
	// {name:张三 age:30} %+v 结构体的键值对都输出了
	// main.User{name:"张三", age:30} 包 类型也输出了


}
func main()  {
	defaultValues()
	scope()
}