package main

import (
	"fmt"
	"math"
)

func main()  {
	var s int8 = 127
	var u uint8 = 1
	fmt.Println(s, u)

	s = s + 1 // -128
	u = u - 2 // 255

	fmt.Println(s, u)
	a := 53 // 十进制
	b := 0700 // 八进制  0开头
	c1 := 0xaabbcc // 十六进制 0x开头
	c2 := 0Xddeeff // 十六进制 0X开头
	fmt.Println(a, b, c1, c2)

	d1 := 0b10000001 // 二进制 0b
	d2 := 0B10000001 // 二进制 0B
	e1 := 0o700 // 八进制 0o
	e2 := 0O700 // 八进制 0O
	fmt.Println(d1, d2, e1, e2)

	// "_"分隔符
	a1 := 5_3_7
	b1 := 0b_1000_0111
	c3 := 0_700
	c4 := 0o_700
	d3 := 0x_ab_cd
	fmt.Println(a1,b1,c3,c4,d3)

	var a2 int = 59
	fmt.Printf("%b\n",a2) // 二进制
	fmt.Printf("%d\n",a2) // 十进制
	fmt.Printf("%o\n",a2) // 八进制
	fmt.Printf("%O\n",a2) // 八进制 0o73
	fmt.Printf("%x\n",a2) // 十六进制 小写
	fmt.Printf("%X\n",a2) // 十六进制 大写 3B

	// 浮点型
	var f float32 = 139.8125
	bits := math.Float32bits(f)
	fmt.Printf("%b\n",bits) // 1000011000010111101000000000000

	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0
	fmt.Println(f1==f2) // true

	// 浮点型 字面值
	var f3 float32 = 3.1415
	var f4 float32 = .15
	var f5 float32 = 81.00
	var f6 float32 = 82.
	fmt.Println(f3, f4, f5, f6)  //3.1415 0.15 81 82

	// 浮点型 科学计数法 十进制
	var f7 float32 = 314.15926e-2
	var f8 float32 = .1234567e+5
	fmt.Println(f7, f8) // 3.1415925 12345.67
	// 浮点型 科学计数法 十六进制

	// 格式化输出浮点型
	var f9 float32 = 3.1415
	fmt.Printf("the float number is %f\n", f9)


}


