package array

import (
	"testing"
)

/*
一 数组的声明
方式1 声明并初始化为零值
var a [3]int

方式2 声明并初始化
b := [3]int{1,2,3}

如果不想写数组的长度 可以用 ... 代替数组的长度，go会默认添加数组的长度
arr2 := [...]int{3,5,7,9}


多维数组的声明与初始化
c := [2][2]int{{1,2},{3,4}}

二 数组的访问和修改
通过下标或索引
访问
arr[下标]
修改
arr[2] = 99

三 数组的遍历
方式1
for i:=0;i<len(arr2);i++

方式2 使用range 关键字 返回的两个是索引 与 值
for idx,e := range arr2

如果不关心索引值,有多个返回值时可以使用占位符_
for _,e := range arr2

四 数组的截取 类似于Python中的切片
a[开始索引(包含),结束索引(不包含)]
非常实用的功能
	t.Log(arr2[0:2])
	t.Log(arr2[2:])
	t.Log(arr2[:3])
	t.Log(arr2[:]) 不填写前后两个值 默认都是截取了

切片 slice
容易造成程序GC的数据结构
 */

func TestArrayInit(t *testing.T){
	var arr1[3]int
	arr2 := [...]int{3,5,7,9}

	t.Log(arr1[0],arr1[2]) // 0,0
	t.Log(arr2)

	arr2[3] = 99
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T){
	//arr2 := [...]int{3,5,7,9}
	//for i:=0;i<len(arr2);i++{
	//	t.Log(arr2[i])
	//}
	//
	//for idx,e := range arr2{
	//	t.Log(idx,e)
	//}
	//
	//for _,e := range arr2{
	//	t.Log(e)
	//}

	arr := [...]int{11,33,55,77,99}
	for i:=0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for idx,v := range arr{
		t.Log(idx,v)
	}

	for _,v := range arr {
		t.Log(v)
	}
}

func TestArraySection(t *testing.T){
	arr2 := [...]int{3,5,7,9}
	t.Log(arr2[0:2])
	t.Log(arr2[2:])
	t.Log(arr2[:3])
	t.Log(arr2[:])
}
//b := [3]int{1,2,3}
//c := [2][2]int{{1,2},{3,4}}