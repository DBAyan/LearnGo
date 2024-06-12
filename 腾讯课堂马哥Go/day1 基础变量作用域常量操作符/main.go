package main

import (
	"fmt"
	"goFirstProgram/mypath" // 不同module下的导入
	//"operator/operator" // 同一个module下的不同包 module_name(模块名称)/dictory_name(文件夹名称)
	zuoyongyu "operator/operator"
	"runtime"
	"strconv"
)

var (
	A = 3 // 大写开头 其他package也可以访问
	b = 4 // 小写开头 仅在本package可以访问
)

func atc()  {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b
	var m int = 8
	var n int = 3
	var g int = m % n
	fmt.Printf("c = %f\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func rel()  {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 5
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a > b && a < c)
	fmt.Println(a > b || a < c)
	fmt.Println(!(a > b))
}

func bitOp()  {
	fmt.Printf("os arch %s,int size %d\n",runtime.GOOS,strconv.IntSize)
	fmt.Printf("260 %b\n", 260)
	fmt.Printf("%b\n", 6|5) // 逻辑或  110 | 101 = 111
	fmt.Printf("%b\n",6 & 5) // 逻辑与 110 & 101 = 100
	fmt.Printf("%b\b", 6 ^ 5) //逻辑 110 ^ 101 = 011
	fmt.Printf("%b\n",^ 5) // 101  1-110
	fmt.Printf("%b\n", 6 >> 2) // 110 11
	fmt.Printf("%b\n", 6 << 2) // 110 11000
}

func foo()  {
	b := 5
	fmt.Printf("A:%d\n", A)
	fmt.Printf("b:%d\n", b)
	{c := 6
	fmt.Printf("b:%d\n", c)
	}
	//fmt.Printf("b:%d\n", c) //  undefined: c
	fmt.Println(zuoyongyu.AA)// 引用的时候是package_name.变量名 或者 package_name.函数名
	res := mypath.Add(111,222)
	fmt.Println(res)

}


func main()  {
	//atc()
	//rel()
	//bitOp()
	//foo()
	zuoyongyu.Scope2()
}
