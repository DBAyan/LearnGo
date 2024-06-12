package main

import "fmt"

// 变量声明方的三种方式
// 一 完整的变量声明方式
// var a int = 10

// var 声明变量的关键字
// a 变量名
// int 变量类型
// 10 变量的初值

func main()  {
	var a int = 10

	// 变量的零值
	var b int
	fmt.Println(a,b)

	// 多个变量声明方式一 ：变量声明快
	var (
		c int = 13
		d string = "hello"
		e int8 = 6
		f rune = 'A' // 65
		g bool = true
	)
	fmt.Println(c,d,e,f,g)
	// 多个变量声明方式二 ：同一行声明同一个类型的多个变量

	var i, j, k int = 1, 2, 3
	var x, y ,z rune = 'a', 'b', 'c' //  97 98 99

	fmt.Println(i,j,k,x,y,z)

	// 声明方式二：省略类型信息的声明
	// var varName = initExpression
	var aa = 12
	fmt.Printf("value %d,type %T\n",aa,aa)

	// 声明方式三：短变量声明
	bb := "hello world"
	cc := true
	dd := 'a' // 97
	ee := 12
	fmt.Println(bb,cc,dd,ee)
}
