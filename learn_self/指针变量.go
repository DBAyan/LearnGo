package main

import "fmt"

//普通变量存储的是值本身，而指针变量存储的是一个地址，这个地址指向存储在内存中的值。我们可以通过在变量名前加上 & 符号来获取变量的地址，称之为取地址操作符。
//指针变量在许多情况下非常有用，特别是当我们需要在函数之间传递大量的数据时。
//通过传递指针而不是实际的数据，可以提高程序的性能，因为数据不需要被复制。
//此外，指针变量还可以用来修改函数外部的变量值，即通过指针间接修改变量的值。

func main()  {
	x := 111
	y := &x //指针变量 y 指向x变量的地址 。& 取地址符
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y) // *y 表示获取指针 y 所指向的变量的值，即 x 的值。
	fmt.Printf("%T\n",y)

	fmt.Println("before:",x)
	midifyValue(&x)
	fmt.Println("after:",x)

}

func midifyValue(p *int)  {
	*p = 20
}