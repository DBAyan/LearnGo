package main

import "fmt"

/*
一 方法本质  method
Go 方法实质上是以方法的 receiver 参数作为第一个参数的普通函数

二 receiver参数类型

值类型   func (t T) M1()  --> func F1(t T)
指针类型 func (t *T) M2() --> func F2(t *T)

当receiver为T时
因为函数参数 值拷贝 ，所以t 是类型T的一个实例副本，在函数中对t的修改不会影响 原类型T
同理 M1方法体中对副本的任何修改操作，都不会影响到原 T 类型实例

当receiver为*T时
会修改原T类型实例

三 选择receiver类型的原则
1：如果 Go 方法要把对 receiver 参数代表的类型实例的修改，反映到原类型实例上，
那么我们应该选择 *T 作为 receiver 参数的类型。

无论是 T 类型实例，还是 *T 类型实例，都既可以调用 receiver 为 T 类型的方法，
也可以调用 receiver 为 *T 类型的方法

2

3 T 类型是否需要实现某个接口


*/

type T1 struct {
	a int
}

func (t T1) M1()  {
	t.a = 10
}

func (t *T1) M2()  {
	t.a = 11
}

func main()  {
	var t T1
	var ptr *T
	fmt.Println(t) // {0}
	println(&t) // 0xc000092f40
	fmt.Println(&t)  // &{0}
	fmt.Println(*&t) // {0} 先取地址
	fmt.Println(ptr) // <nil>
	fmt.Println()


	println(t.a) // 0

	t.M1()
	println(t.a) // 0

	t.M2()
	println(t.a) // 10

	var t2 = &T1{}
	println(t2.a)
	t2.M1()
	println(t2.a)
	t2.M2()
	println(t2.a)
}