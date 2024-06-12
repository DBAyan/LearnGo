package main

import (
	"fmt"
)

/*
error is value
error type

一 sentinel error
定义：
预定义的特定的某种错误

格式：
if err == errSomething{...}

缺点 ：
1 不灵活的错误处理方式
fmt.Errorf 携带上下文的时候，也会破坏调用者的 ==

2 Sentinel Error 会成为你API的公共部分 会增加你API暴雷的表面积
 接口表面积越大 接口月脆弱

3 Sentinel Error在两个包之间创建了 依赖

结论 尽可能避免使用Sentinel Error，不建议使用 模仿

推荐的方式
二 Error types
定义：
Error type是实现了 error 接口的自定义类型（struct）
改进 ：
提供更多的上下文


例子:
os.PathError

需要学习 类型断言 ， 类型Switch，空接口 等
类型断言
https://blog.csdn.net/General_zy/article/details/124572304

switch err := err.(type)

三 Opaque Errors 不透明错误处理
灵活
调用者 与 源码 耦合最少的。

*/


type myError struct {
	Msg string
	File string
	Line int
}

func (e *myError) Error()  string {
	return fmt.Sprintf("%s:%d: %s",e.File,e.Line,e.Msg)
}

func test() error {
	return &myError{"some error happened","Server.go",42}
}

func main()  {
	err := test()
	switch err := err.(type) {
	case nil:
	case *myError:
		fmt.Println(err.Error())
	default:

	}
	// switch type 任何类型都实现了空接口
	var x interface{} = func() {} // "yhh",12,3.14
	switch i:=x.(type) {
	case nil:
		fmt.Println(i,"type nil!")
	case int:
		fmt.Println(i,"type int")
	case string:
		fmt.Println(i,"type string")
	case bool:
		fmt.Println(i,"type bool")
	case float64:
		fmt.Println(i,"type float64")
	default:
		fmt.Println("undefined")
	}
}