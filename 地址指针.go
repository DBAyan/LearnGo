package main

import "fmt"

func swap(a,b int)  {
	//t := *a
	//*a = *b
	//*b = t
	t := a
	a = b
	b = t
}

func main()  {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p,%p\n",&cat,&str) // 0xc0000b2008,0xc000096210
	var house = "ci qu jing yuan nan qu" // 字符串类型
	ptr := &house // & 取这个变量的地址
	fmt.Printf("%T\n",ptr) // ptr的变量类型  *string 字符串的指针类型

	fmt.Printf("%p\n",ptr) // ptr变量的内存地址 0xc000010240 0x 表示16进制

	value := *ptr // 对指针进行取值操作

	fmt.Printf("%T\n",value) // string

	fmt.Printf("%s\n",value) // ci qu jing yuan nan qu

	x, y := 11,22
	swap(x,y)
	fmt.Println(x,y)

	aa := 111
	fmt.Println(&aa)  // 0xc000016078
	bb := 222
	fmt.Println(&bb) // 0xc000016090

	t := aa
	aa = bb
	bb = t
	fmt.Println(&bb) // 0xc000016090

	fmt.Println(&aa) // 0xc000016078

	fmt.Printf("a : %d,b:%d",aa,bb)

}