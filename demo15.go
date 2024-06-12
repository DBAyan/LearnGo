package main

import "fmt"

func main(){

	s1 := make([]int,5)
	fmt.Printf("The length od s1: %d.\n", len(s1))
	fmt.Printf("The capacity of s1 : %d.\n",cap(s1))
	fmt.Printf("The value of s1: %d.\n",s1)

	s2 := make([]int,5,8)
	fmt.Printf("The length od s2: %d.\n", len(s2))
	fmt.Printf("The capacity of s2 : %d.\n",cap(s2))
	fmt.Printf("The value of s2: %d.\n",s2)

	s3 := []int{1,2,3,4,5,6,7,8}
	s4 := s3[3:6]
	fmt.Printf("The length od s4: %d.\n", len(s4))
	fmt.Printf("The capacity of s4 : %d.\n",cap(s4))
	fmt.Printf("The value of s4: %d.\n",s4)

	// 初始化数组/定义数组

	// 方法1
	//
	var a = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(a) // [1000 2 3.4 7 50]

	// 方法2
	a2 := [2]int{0,9}
	fmt.Println(a2) //[0 9]


}
