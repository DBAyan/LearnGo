package main

import (
	"fmt"
	"unsafe"
)

// bool
// string
// 数字类型
	// 有符号整型 int（建议使用） int8 int16 int32 int64
	// 无符号整形 uint uint8 uint16 uint32 uint64
	// 浮点型 float32 float64
	// 复数类型 complex64 complex128
	// byte uint8的别名
	// rune int32的别名


func main(){
	a := true
	b := false
	fmt.Println("a:" ,a,"b:",b)
	var c,d bool
	c = a && b
	d = a || b
	fmt.Println(c,d)

	var e int =89
	f:=3
	fmt.Println("value of e is", e, "and f is", f)
	fmt.Printf("Type of e is %T，size of e is %d\n",e,unsafe.Sizeof(e))
	fmt.Printf("Type of e is %T，size of e is %d",f,unsafe.Sizeof(f))


}
