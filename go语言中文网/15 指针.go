package main

import "fmt"

func main() {
	b := 255
	// 指针类型的声明 *T *int
	// &b 取地址 取变量b的地址赋值给a,a是指针类型
	var a *int = &b
	fmt.Printf("type of a is %T\n",a)
	fmt.Println("addres of b is ",a)

	// 指针的零值为nil
	var c *string
	fmt.Println(c)

	// 指针的解引用 *e
	d := "abcd"
	e := &d
	fmt.Println("address of d is",e)
	fmt.Println("value of d is",*e)

	// 向函数传递指针参数
	aa := 5588
	fmt.Println("value of a before function call is",aa)
	bb := &aa
	change(bb)
	fmt.Println("value of a before function call is",aa)
}
func change(val *int){
	*val = 55
}