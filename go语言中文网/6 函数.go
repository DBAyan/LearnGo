package main

import "fmt"

// 1 函数定义
// 2 函数声明
/*
func functionname(parametername type) returntype{
 函数体（具体实现的功能）
}
 */

//3 参数列表 与 返回值非必须的 ,以下的函数声明也是有效的
//func functionname(){}

func calculateBill(price int,no int) int {
	var totalPrice = price * no
	return totalPrice
}

//4 参数类型如果相同 可以 如下简写
//(price ,no int)

//5 调用函数
// functionname(实参)

//6 多返回值
// 计算矩形的面积 和 周长

//func rectProps(length,width float64)(float64, float64){
//	var area = length * width
//	var perimeter = (length + width) * 2
//	//return area, perimeter
//	return area, perimeter
//}
// 7 命名返回值

func rectProps(length,width float64)(area, perimeter float64){
	area = length * width
	perimeter = (length + width) * 2
	return //不需要明确指定返回值，默认返回 area, perimeter的值
}

func main(){
	price, no := 90 ,6
	totalPrice := calculateBill(price, no)
	fmt.Println("total price is", totalPrice)
	area, perimeter := rectProps(10.8,5.6)
	fmt.Printf("Area %f Perimeter %f \n", area, perimeter)
//	值需要面积 不需要周长
	area2, _ := rectProps(111,456)
	fmt.Printf("Area2 %f", area2)
}


// 8 空白符 _
// _ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。

