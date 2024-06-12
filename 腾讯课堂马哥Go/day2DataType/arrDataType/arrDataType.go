package main

import "fmt"

// 传值
func arrPiont(arr [5]int)  {
	fmt.Printf("arr point is %p\n", &arr)
	arr[0] = 10
	fmt.Printf("arr point is %p\n", &arr)
}

// 传地址
func arrPoint2(arr *[5]int){
	fmt.Printf("arr point is %p\n", arr)
	arr[0] = 10
	fmt.Printf("arr point is %p\n", arr)
	fmt.Printf("")
}

func main()  {
	// go 声明变量的几种方式  https://blog.csdn.net/weixin_62281810/article/details/124574138
	// 数组初始化
	// var 变量名 [数组长度]数据类型 = [数组长度]数据类型{}
	var arr1 [5]int = [5]int{} // 数组必须制定长度和类型，且长度和类型指定后不可改变
	var arr2 = [5]string{} // 不指定数据类型 ，根据类型推导
	var arr3 = [5]int{11, 22} // 给前两个元素赋值
	var arr4 = [3]string{0:"yan", 2:"hang"} // 给第一个 和 第三个元素赋值
	var arr5 = [...]int{0,1,2,3} // 不指定数组长度 ，根据元素个数推断数组长度

	var arr6 = [...]struct{
		name string
		age int} {{"nico",29}, {"liffy",30}}

	// 二维数组的初始化
	// 5行 3列
	var arr7 = [5][3]int{{1,2,3}, {11,22,33}, {111,222}} //  [[1 2 3] [11 22 33] [111 222 333] [0 0 0] [0 0 0]]

	var arr8 = [...][2]string{{"a","b"},{"aa","bb"},{}}

	fmt.Println(arr1)
	fmt.Printf("type arr1:%T\n",arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)

	// 访问数组的元素
	var arr [5]string = [5]string{"a","b","c","d", "e"}
	fmt.Printf("arr first ele is %s\n", arr[0]) // 第一个元素
	fmt.Printf("arr last ele is %s\n", arr[len(arr) - 1]) // 最后一个元素

	// 访问多维数组
	fmt.Println(arr7[1][1]) // 第二行的第二个元素

	// 修改数组的元素
	arr[0] = "didi"
	fmt.Println(arr)

	// 遍历数组
	for i, ele := range arr{
		fmt.Printf("arr %d ele is %s\n",i,ele)
	}
	// 遍历数组
	for i:=0;i<len(arr); i++ {
		fmt.Printf("arr the %d's ele is %s\n",i,arr[i])
	}
	// 遍历二位数组
	for row, array := range arr7 {
		fmt.Printf("%d row is %v\n",row,array)
		for idx ,eles := range array {
			fmt.Printf("row %d colum %d（arr7[%d][%d]） : eles %d\n",row, idx,row, idx,eles)
		}
	}
	// cap len
	fmt.Printf("arr7 length is %d, arr7 capacity is %d\n", len(arr7),cap(arr7))

	// 数组作为函数参数 传值，地址是不同的
	fmt.Printf("arr1 point is %p\n", &arr1)
	arrPiont(arr1)
	arrPoint2(&arr1)

}
