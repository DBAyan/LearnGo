package main

import "fmt"

// 什么是变量


func main() {
	// 方法一 只声明变量
	// var name type
	// 如果变量未被赋值，Go 会自动地将其初始化，赋值该变量类型的零值（Zero Value）
	var age int
	fmt.Printf("my age is %d\n", age) // my age is %d 0
	age = 29
	fmt.Printf("my age is %d\n",age)

	//方法二 声明并赋值
	//var name type = initialvalue
	var name string = "yhh"
	fmt.Printf("my name is %s\n", name)

	// 方式三 省去变量类型 通过类型推断
	// var name = initialvalue
	var habbit = "play game"
	fmt.Printf("i like %s \n" ,habbit)

	// 方式四 简短声明 变量名 := 变量值
	// name := initialvalue
	salary := 35000
	fmt.Printf("My salary is %d\n",salary)

	// 声明多个变量
	var width,height int = 100,50
	fmt.Printf("width is %d,height is %d", width,height)

}