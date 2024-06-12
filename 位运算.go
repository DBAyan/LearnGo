package main

import "fmt"

// https://www.runoob.com/go/go-operators.html

func main()  {
	var a int8 = 60
	var b int8 = 13
	var c int8
	fmt.Printf("%b\n", a) // 0011 1100

	fmt.Printf("%b\n", b) // 0000 1101
	// 按位与 都为1 才为1
	c = a & b
	fmt.Printf("%b\n",c) // 0000 1100
	// 按位或 有一个为1 便为1
	c = a | b
	fmt.Printf("%b\n",c) // 0011 1101
	// 按位异 相同为0  不同为1
	c = a ^ b
	fmt.Printf("%b\n",c) // 0011 0001
	// 左移运算符 高位丢弃 低位补0
	c = a << 1
	fmt.Printf("%b\n",c) // 0111 1000

	// 右移元素安抚 高位补0 低位丢弃
	c = a >> 1
	fmt.Printf("%b\n",c) // 0001 1110

}
