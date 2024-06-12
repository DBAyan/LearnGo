package main

import (
	"fmt"
)

// TestSliceGrowth append 切片扩容

func TestSliceGrowth()  {
	var s []int
	for i:=0;i<10;i++ {
		s = append(s, i)
		fmt.Printf("len:%d,cap:%d\n",len(s),cap(s))
	}

}


// 没有按照规律扩容 内存对齐

func TestSliceCap()  {
	s := make([]int,0,70)
	prevCap := cap(s)
	for i := 0; i<1300; i++ {
		s = append(s, 0)
		currCap := cap(s)
		if currCap > prevCap{
			fmt.Printf("cap %d --> %d\n",prevCap,currCap)
			prevCap = currCap
		}
	}
}

func subSlice()  {
	arr := make([]int, 5, 5)
	crr := arr[0:2] // 左包右不包
	crr[1] = 8 // 子切片修改元素 ，母切片也会被修改
	fmt.Printf("arr:%v,point:%p\ncrr:%v,point:%p\n", arr,arr, crr,crr)

	// 内存分离后 首元素的地址不再相同
	crr = append(crr,9,10,11,12)
	fmt.Printf("arr:%v,point:%p\ncrr:%v,point:%p\n", arr,arr, crr,crr)

}

// 切片作为函数参数时
func updateSlice(sli []int)  {
	sli[0] = 111
	fmt.Printf("函数内 sli:%v\n",sli)
}

func main()  {
	// 切片的声明与初始化
	var s []int  //切片声明 len=cap=0
	s = []int{} // 初始化 len=cap=0

	// 工作中常用的方式make
	s1 := make([]int, 3) // 声明并初始化 len=cap=3
	s2 := make([]int, 3, 5) // 声明并初始化 len=3 cap=5
	s3 := []int{1,2,3,4,5} // 声明并初始化 len=cap=5

	// 二维数组
	s2d := [][]int{
		{1},{2,3}, // 二维数组各行的列数相等，但二位切片各行的len可以不等
	}

	fmt.Printf("s:%v,length: %d, capacity :%d\n",s,len(s),cap(s))
	fmt.Printf("s1:%v,length: %d, capacity :%d\n",s1,len(s1),cap(s1))
	fmt.Printf("s2:%v,length: %d, capacity :%d\n",s2,len(s2),cap(s2))
	fmt.Printf("s3:%v,length: %d, capacity :%d\n",s3,len(s3),cap(s3))
	for row,arr := range s2d{
		fmt.Printf("the %d row ,arrry:%v, length %d\n", row,arr, len(arr))
	}
	// append
	s3 = append(s3,100,200)
	fmt.Printf("s3:%v,length: %d, capacity :%d\n",s3,len(s3),cap(s3))

	//TestSliceGrowth()

	//TestSliceCap()

	//subSlice()
	sli := make([]int,1,1)
	fmt.Printf("函数外修改之前 sli:%v\n",sli)

	updateSlice(sli)
	fmt.Printf("函数外 sli:%v\n",sli)

}
