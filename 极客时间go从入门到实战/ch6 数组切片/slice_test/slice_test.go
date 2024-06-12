package slice_test

import "testing"

/*
切片 比较复杂 容易造成GC的数据结构

一  切片的声明
方式1 不用指定长度 ，是可变长
var 变量名 []类型
var s0 []int

方式2 声明并初始化
s1 := []int{1,2,3}

方式3 使用make关键字
三个参数 len cap
s2 := make([]int,3,5)

capcity 代表容量
length 代表可访问元素的个数


数组 VS 切片
1 容量是否可伸缩 是否可变长
切片的容量可伸缩
与数组不同的是 没有指定长度 slice 是可变长的

2 是否可以比较
数组可比较 相同维度 相同长度 每个元素都相同 则两个数组相同

切片不可比较 invalid operation: s0 == s1 (slice can only be compared to nil)
只能和nil做比较


二 切片是如何实现可变长的
当切片中的元素个数 len 要超过 容量 cap时，容量是*2的增长。


* 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
* 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
* 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
* 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）

添加元素为什么要写成这样？
s = append(s,i)
结构体指向的连续存储空间地址发生了变化 ，所以并不总是往这个连续的存储空间中追加值 ，
当存储空间扩展的时候，他会创建一个新的存储空间并且把原有的数值都copy到新的连续存储空间中，
所以slice使用很方便，不用像数组一样初始化的时候预估容量大小，可以自增长。
但是slice自增长也有代价，比如存储空间的复制，数据的拷贝等，后面性能调优会讲到。


三  切片共享存储结构
切片是一个结构体 ，指向后端连续的存储空间地址。
我们有多个slice 指向同一个array

四 切片的遍历 同数组

*/

func TestSliceInit(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))

	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3}
	t.Log(len(s1),cap(s1))

	s2 := make([]int,3,5)
	t.Log(len(s2),cap(s2))
	//t.Log(s2[0],s2[1],s2[2],s2[3]) //index out of range [3] with length 3
	t.Log(s2[0],s2[1],s2[2])

	s2 = append(s2,10)
	t.Log(len(s2),cap(s2))
	t.Log(s2[0],s2[1],s2[2],s2[3])

}

func TestSliceGrowing(t *testing.T){
	var s []int
	for i:=0;i<10;i++{
		s = append(s,i)
		t.Log(len(s),cap(s))
	}
}


func TestSliceSharedMemory(t *testing.T){
	year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}

	Q2 := year[3:6]
	t.Log(Q2,len(Q2),cap(Q2)) //  [Apr May Jun]  3  9

	summer := year[5:8]
	t.Log(summer,len(summer),cap(summer))  //  [Jun Jul Aug] 3 7

	summer[0] = "unknow" // 修改summer
	t.Log(Q2) // 输出 Q2   [Apr May unknow]
	t.Log(year) // [Jan Feb Mar Apr May unknow Jul Aug Sep Oct Nov Dec]
}

func TestSliceCompare(t *testing.T){
	s0 := []int{1,2,3}
	//s1 := []int{3,5,4}
	//t.Log(s0 == s1) //  invalid operation: s0 == s1 (slice can only be compared to nil)
	t.Log(s0 == nil)
}

func TestSliceLoop(t *testing.T)  {
	s0 := []int{1,22,333}
	for idx,v := range s0{
		t.Log(idx,v)
	}
}