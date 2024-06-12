package main

import (
	"fmt"
	"goFirstProgram/mypath"
	"gonum.org/v1/gonum/stat"
)

// 第一个go 程序 hello world
// 第二个go 程序 跨包引用
// 第三个go 程序 第三方包
// https://blog.csdn.net/weixin_46022274/article/details/124706476
/*
go get -u gonum.org/v1/gonum/stat
-u 参数 表示如果本地已经下载了这个包会更新, 如果没有就下载

*/

// 函数入口 必须满足在package  main 下的main 函数
// VS code  可以设置在ctrl + s 保存的时候格式化

func main() {
	fmt.Println("hello world")
	a := 3
	b := 4
	c := mypath.Add(a, b)
	fmt.Println(c)
	arr := []float64{1, 2, 3, 4, 5.5}
	v := stat.Variance(arr, nil)
	fmt.Printf("方差=%f\n", v)

}
