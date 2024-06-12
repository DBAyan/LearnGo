package main

import "fmt"

// 方式一 声明并且赋0值
var x int
// 方式二 声明并且赋值
var y int = 11
// 省略类型
var z = 12

var f float32 = 3.14

type Person struct {
	name string
	age int
}


var s  string = "我是字符串"


func main()  {
	x = 1
	// 方式三 短变量声明。短变量声明是一种快速定义和初始化变量的方法，可以在函数内部使用，但不能在函数外部使用。
	m := 2
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(m)
	fmt.Println(s)

	//结构体初始化 方式一 字面量初始化
	//可以使用字面量初始化语法初始化结构体，类似于C++中的方法。在声明结构体变量时，可以使用花括号和逗号分隔的值列表来初始化结构体的各个字段
	p := Person{"yanhaihang",32}
	fmt.Println(p)

	//结构体初始化 方式二 使用结构体字面量初始化
	//可以直接在定义结构体变量时初始化其字段。
	var p2 Person
	p2.age = 25
	p2.name = "nico"
	fmt.Println(p2)

	//结构体初始化 方式三 使用new函数进行结构体初始化
	//在Go中，可以使用new函数来分配并初始化一个新的结构体实例，并返回一个指向该实例的指针。例如：
	p3 := new(Person)
	p3.age= 24
	p3.name = "xiaojuan"
	fmt.Println(p3)
	fmt.Printf("%T",p3)
}

