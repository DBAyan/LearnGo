package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
一 声明方式
func 函数名 (参数列表) (返回值列表) {
	函数体
}
1 func 关键字
2 函数名
3 参数列表 (参数名1 参数类型,参数名2 参数类型) 支持变长类型参数
4 返回值列表 具名返回值
5 函数体

二 横向对比变量声明
函数签名

s := T{}
f := func(){}

匿名函数

三 函数参数
形式参数 形参 函数声明阶段
实际参数 实参 函数调用阶段

函数参数财传递 值拷贝
当参数为
1 整型 数组 结构体 传递的是内存中的实际内容
2 字符串 切片 map 传递的是描述符
3 接口类型 编译器 会把 传递的实参赋值给对应的接口类型形参
4 变长类型参数  Go 编译器会将零个或多个实参按一定形式转换为对应的变长形参。
变长参数在内部是通过切片实现的

四 多返回值
函数支持多返回值
1 返回值形式
a 无返回值形式 func foo() {}
b 1个返回值形式 func foo() error {}
c 多个返回值形式 func foo() (int,string,error) {}

2 具名返回值

3 最佳实践 如何选择具名返回值

五 函数 一等公民
特征一： go函数可以存放在变量中
特征二：可以在函数内创建函数 并通过返回值返回
闭包
特征三：作为参数传入函数
特征四 ：拥有自己的类型

六 函数的妙用
1 函数也可以显式转型
2 利用闭包简化函数调用


*/

func myAppend (sl []int, elems ...int) []int {
	fmt.Printf("%T\n",elems) // []int 变长参数在内部是通过切片实现的
	if len(elems) == 0 {
		fmt.Println("no elems to append")
		return sl
	}
	sl = append(sl,elems...)
	return sl
}

func setup(task string) func()  {
	println("do some setup stuff for",task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func greeting(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w, "welcome,Gopher!\n")
}

func times(x,y int)  int {
	return x*y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x,y)
	}
}


func main()  {
	sl := []int{1,2,3}
	sl = myAppend(sl)
	fmt.Println(sl)
	sl = myAppend(sl,4,5,6)
	fmt.Println(sl)

	var myPrintF = func(w io.Writer,format string,a ...interface{}) (int,error) {
		return fmt.Fprintf(w, format, a...)
	}

	fmt.Printf("%T\n",myPrintF) // func(io.Writer, string, ...interface {}) (int, error)
	myPrintF(os.Stdout,"%s\n","Hello Go") // Hello Go

	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")

	//time.AfterFunc(time.Second*2, func() { println("timer fired") })
	// http.ListenAndServe(":8080",http.HandlerFunc(greeting))

	println(times(2,5))
	times(4,5)
	times(7,5)

}