package main

import "fmt"

// 什么是常量
// 常量的特点
func main()  {
	type myInt int
	const n myInt = 15
	// const m int = n + 3 // cannot use n + 3 (type myInt) as type int in const initializer
	// var a int = 13
	//fmt.Println(a + n) // invalid operation: a + n (mismatched types int and myInt)
	const c = 12 // 无类型常量
	fmt.Println(n + c)

	// 用常量枚举类型
	// 特性一 代码自动重复上一行
	const(
		Apple, Bananer = 11, 22
		Strawberry, Grage // 等价于 Strawberry, Grage = 11, 22
		Peer, Watermelon //  等价于 Peer, Watermelon = 11, 22
	)
	fmt.Println(Apple, Bananer, Strawberry, Grage, Peer, Watermelon) // 11 22 11 22 11 22
	// << 位运算 二进制右移 在最右边加0
	// 1的二进制为  0000 0001
	// 1<<2  后  二进制为 0000 0100 十进制为 4
	//
	const (
		mutexLocked = 1 << iota // 1<<0 iota为0
		mutexWoken // iota为1 1<<1 2
		mutexStarving // iota为2 1<<2 4
		mutexWaiterShift = iota // iota为3 3
		starvationThresholdNs = 1e6
	)
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs) // 1 2 4 3 1e+06
	const (
		aa = iota + 1 // iota:0 1
		bb // iota:1 2
		cc // iota:2 3
	)
	const (
		i = 1 << iota // iota:0 1
		j // iota:1 2
		k // iota:2 4
 	)
	fmt.Println(aa, bb, cc,i,j,k)
	const (
		_ = iota // 0
		IPV6_V6ONLY // 1
		SOMAXCONN // 2
		SO_ERROR // 3
	)
	fmt.Println(IPV6_V6ONLY, SOMAXCONN, SO_ERROR)

	const (
		_ = iota
		pin1
		pin2
		pin3
		_
		pin5
	)
	fmt.Println(pin1, pin2,pin3,pin5)
}
