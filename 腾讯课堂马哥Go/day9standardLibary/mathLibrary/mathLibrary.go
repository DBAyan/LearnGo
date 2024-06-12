package main

import (
	"fmt"
	"math"
)

// 数学库中的常量

func constant(){
	fmt.Println(math.E)  // 自然数 2.718281828459045
	fmt.Println(math.Pi) // 圆周率 3.141592653589793
	fmt.Println(math.MaxInt) // 最大的整数 9223372036854775807
	fmt.Println(math.MinInt) // 最小的整数-9223372036854775808
	fmt.Println(math.SmallestNonzeroFloat64)
	fmt.Println(math.SmallestNonzeroFloat32)

}

func mathFunc()  {
	//取整
	fmt.Println(math.Ceil(2.5)) // 向上取整 3
	fmt.Println(math.Floor(2.5)) // 2
	fmt.Println(math.Ceil(-3.4)) //  向上取整 -3
	fmt.Println(math.Floor(-3.4)) // -4

	// 截取整数部分
	fmt.Println(math.Trunc(5.6)) // 取整数部分 5
	fmt.Println(math.Trunc(5.4)) // 5

	// 取绝对值
	fmt.Println(math.Abs(-1.2)) // 取绝对值

	// 取绝对值
	fmt.Println(math.Modf(1.2))  // 1.2

	// 取两个数的最大值或者最小值
	fmt.Println(math.Max(2,5)) // 取两个数的最大值 5
	fmt.Println(math.Min(2,5)) // 去两个输的最小值 2

	//  返回差值
	fmt.Println(math.Dim(6,2))

	// 平方

	//根号下

	// LOG
	//随机数

	//打乱

	// 三角函数


}

func main()  {
	constant()

	mathFunc()
}
