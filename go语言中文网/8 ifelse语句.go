package main

import "fmt"

/*

语法1
if condition {

}
condition为真 执行 {}之间的代码

语法2 配合 else 和 else if

if condition{
} else if condition {
} else {
}
可以有多个 else if 语句分支

语法3
if statement;condition{
}
注意 ：
else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。


*/

//func main (){
//	//num := 10
//	if num := 10; num%2 != 0 {
//		fmt.Println("the num is odd")
//	} else {
//		fmt.Println("the num is even")
//	}
//}


func main(){
	num := 20
	if num % 2 == 0{
		fmt.Printf("the num %d is even",num)
	}else {
		fmt.Printf("the num %d is odd",num)
	}
	if num2 := 5;num2 % 2 == 0 {
		fmt.Printf("the number %d is even",num2)
	} else {
		fmt.Printf("the number %d is odd",num2)
	}

}