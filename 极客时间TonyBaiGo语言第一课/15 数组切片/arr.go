package main

import (
	"fmt"
	"unsafe"
)

func foo(arr [5]int)  {

}

func main()  {
	var arr1 [5]int
	// var arr2 [6]int
	// var arr3 [5]string
	foo(arr1)
	// 长度与类型都要相等数组才相等
	// foo(arr2) //  cannot use arr2 (type [6]int) as type [5]int in argument to foo
	// foo(arr3) //  cannot use arr3 (type [5]string) as type [5]int in argument to foo

	// 数组的声明方式一 ：声明变量并初始化
	var arr = [6]int{1,2,3,4,5,6}
	fmt.Printf("The length of arr is %d.\n", len(arr)) // 元素数
	fmt.Printf("The bytes of arr is %d.\n",unsafe.Sizeof(arr)) // byte 字节数 int8 6*8 = 48
	// 声明变量 初始值为对应类型的零值
	var arr3  [3]string
	fmt.Println(arr3) // [  ]

	// 数组的声明方式二 ：短变量声明
	arr4 := [...]string{"yan","hai","hang"}
	fmt.Printf("The length of arr is %d.\n", len(arr4))

	// 通过下标访问数组中的元素
	var arr5 = [...]int{99:39}
	fmt.Println(arr5)

	var arr6 = [7]int{1,2,3,4,5,6,7}
	//fmt.Println(arr6[8]) // invalid array index 8 (out of bounds for 7-element array)
	//fmt.Println(arr[-1]) //  invalid array index -1 (index must be non-negative)
	fmt.Println(arr6[0],arr6[5])

	// 多维数组
	var mArr = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(mArr)

	s1 := make([]byte,8,10)
	fmt.Println(s1)

	// 切片 slice
	var nums = []int{2,4,6,8}

	fmt.Println(len(nums),cap(nums)) // 4,4
	nums = append(nums,10)
	fmt.Println(len(nums),cap(nums)) // 5,8
}