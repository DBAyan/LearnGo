package main

import (
	"fmt"
	"log"
	_ "math" // imported and not used: "math"
	//"rectangle"
)

/*
1 包的定义
2 用处 包用于组织 Go 源代码，提供了更好的可重用性与可读性
3 导入路径：自定义包 相对于工作区内 src 文件夹的相对路径
使用 go env 查看gopath 为 GOPATH="/Users/yanhaihang/Desktop/go/LearnGo"
4 可以被导出的函数的首字母必须大写
5 init函数  声明 func init(){}
6 init 函数作用  a 可用于执行初始化任务，b 也可用于在开始执行之前验证程序的正确性。
7 包的初始化顺序 先初始化被导入的包 ，然后初始化包级别的变量 ，然后main包里的init函数，
8 空白标识符 Blank identifier _
不使用导入的包 程序会报错。imported and not used:
这样做的目的：防止导入过多的包，减少编译时间。

*/



var rectLen, rectWidth float64 = 6, 7

func init(){
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}
func main(){
	fmt.Printf("Geometrical shape properties")
	//fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	//fmt.Printf("diagonal of rectangle %.2f ",rectangle.Diagonal(rectLen,rectWidth))
}
