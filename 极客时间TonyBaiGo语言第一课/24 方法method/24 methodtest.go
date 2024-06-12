package main

import "fmt"

/*
组合设计哲学

一 方法声明
func (t *T或T) MethodName(参数列表) (返回值列表) {
    // 方法体
}
6个部分：
func
receiver参数
方法名
参数列表
返回值列表
方法体

二 receiver参数

1 作用：
这个 receiver 参数也是方法与类型之间的纽带，也是方法与函数的最大不同。

2 作用域
receiver参数作用域 在方法作用域内唯一 ，不能与 形参名 或者 具名返回值 同名

3 注意：
（1）receiver 参数的基类型本身不能为 指针类型 或 接口类型

（2）方法声明要与 receiver 参数的基类型声明放在同一个包内
	推论：
	a 不能为原生类型添加方法  int string
	b 不能跨越go包为其他包声明方法
(3)方法 M 是类型 T 的方法，也可以通过 *T 类型变量也可以调用 M 方法

三 方法的本质

本质：方法的本质是函数 ，receiver参数是传入函数的第一个参数



*/

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t T)Put(a int) int {
	t.a = a
	return t.a
}



func main()  {
	var tt = T{
		2,
	}

	tt.Get()

	//tt.Put(12)
	fmt.Println(tt.Put(12))
}

