package fibo

import (
	"fmt"
	"testing"
)

// 1 1 2 3 5 8

func TestFibo(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a int = 1
	//	b int = 1
	//)

	// 赋值可以进行类型推断
	a := 1
	b := 1

	fmt.Print(" ",a)
	for i:=0;i<5;i++{
		fmt.Print(" ",b)
		tmp := a
		a = b
		b = tmp + a

	}
}

func TestExchange(t *testing.T){
	a := 2
	b := 3
	//tmp := a
	//a = b
	//b = tmp

	// 在一个赋值语句中可以对多个变量同时进行赋值

	a,b = b,a
	t.Log(a,b)
}