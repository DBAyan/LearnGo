package main

import "fmt"

/*
1 数组 array
数组是同一类型元素的集合。只能是同一类型。

2 声明
var name [n]T
n 代表数组中元素的数量 T代表数组元素的类型

3 数组如果不被赋值 ，数组中的所有元素都被自动赋值为数组类型的零值。

4 数组长度 len()

5 使用range迭代数组

*/

func main (){
	var a [3]int
	fmt.Println(a) // [0 0 0]
	fmt.Println("The length of a is",len(a))

	// 声明时忽略数据长度 ，使用 ... 代替长度，自动计算出长度
	s := [...]string{"yhh","handsome","work","study"}
	fmt.Println(s)
	fmt.Println(len(s))

	// 数据是值类型
	k := s
	k[0] = "zhangsan"
	fmt.Println("s is" , s)
	fmt.Println("k is ", k)

	b := [3]int{756, 34, 87}
	fmt.Println(b) //[1 2 3]
	for i := 0 ;i < len(b); i++ {
		fmt.Printf("%d th element of b is %v \n", i ,b[i])
	}

	c := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for n , v := range c{
		fmt.Printf("%d th element of c is %.2f\n",n,v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)

	d := [5]int{1,2,3,4,5}
	var e []int = d[1:3]
	fmt.Println(d)
	fmt.Println(e)

}