package main

import (
	"fmt"
)

func add(n int) int {
	return n+1
}

func sub(n int) int{
	return n - 1
}

func fall(age int)  {
	switch  {
	case age > 50:
		fmt.Println("当总统")
		fallthrough
	case age > 25:
		fmt.Println("生孩子")
		fallthrough
	case age > 22:
		fmt.Println("结婚")
		fallthrough
	case age > 6:
		fmt.Println("上学")
	}
}
func main()  {
	// switch语句 红绿灯
	//switch - case - default 可以模拟 if - else if - else 但是只能是相等判断
	//switch 和 case后面可以跟常量 变量或函数表达式，只要他们表示的数据类型相同即可
	//case后面可以跟多个值，只要有一个值满足即可

	color := "black"

	switch color {
	case "rea": // 相当于 else if color == "red"
		fmt.Println("stop")
	case "blue":
		fmt.Println("go")
	case "yellow":
		fmt.Println("wait")
	default: // 相当于 else
		fmt.Println("invaild color")
	}

	// switch  与  case  后可以跟函数 ，只要保证数据类型相同即可
	// case 后面可以跟多个值 只要其中一个满足即可
	const B = 8
	switch add(5) {
	case sub(9):
		fmt.Println("sum is 8")
	case 7:
		fmt.Println("sum is 7")
	case 5,3,B:
		fmt.Println("sum is 7")
	case 2+4:
		fmt.Println("sum is 6")
	default:
		fmt.Println("unkown")
	}

	//空 switch

	//switch后面有表达式时，只能模拟相等的清空
	//当switch后面没有表达式时，case就可以跟任意表达式。

	//switch {
	//case 5 > 8 :
	//	fmt.Println(" 5 > 8")
	//case 2 == 4-1:
	//	fmt.Println("2 == 4-1")
	//case true:
	//	fmt.Println("true")
	//default:
	//	fmt.Println("2 == 4-1")
	//}

	// fallthrough 当case 满足条件后也会继续往后执行
	fall(30)

	// switch type
	var num interface{} = 10
	switch value:= num.(type) {
	case int:
		fmt.Printf("num is int %d\n", value)
	case float64:
		fmt.Printf("num is float64  %f\n", value)
	case byte, string, bool:
		fmt.Printf("num is type byte,string,bool %v\n", value)
	default:
		fmt.Printf("noting")
	}
}