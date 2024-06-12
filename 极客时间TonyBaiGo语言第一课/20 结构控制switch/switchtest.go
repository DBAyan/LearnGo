package main

import (
	"fmt"
)
/*
switch 的执行顺序
自上而下
自左向右

*/
func case1() int {
	fmt.Println("eval case1 expr")
	return 1
}

func case2_1() int {
	fmt.Println("eval case2_1 expr")
	return 0
}

func case2_2() int {
	fmt.Println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("eval switch expr")
	return 2
}

func main()  {
	switch switchexpr() {
	case case1():
		println("exec case1")
	case case2_1(),case2_2():
		println("exec case2")
		fallthrough
	case case3():
		println("exec case3")
		fallthrough
	default:
		println("exec default")
		//fallthrough // cannot fallthrough final case in switch
	}


	// type switch

	var x interface{} = false
	switch v := x.(type) {
	case nil:
		fmt.Printf("v is nil")
	case int:
		fmt.Printf("v:%d is int", v)
	case string:
		fmt.Printf("v:%s is string",v)
	case bool:
		fmt.Printf("v:%b is bool" ,v)
	default:
		fmt.Printf("v is not support type")

	}
}