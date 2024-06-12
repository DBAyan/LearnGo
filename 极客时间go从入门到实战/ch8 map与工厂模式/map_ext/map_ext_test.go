package map_ext

import "testing"

/*
map 与 工厂模式
工厂模式是最常用的设计模式之一

map的value 不仅可以是一个值 ，也可以是一个方法
与go的 duck type 接口方式一起 ，可以方便的实现单一方法的对象的工厂模式

使用map实现set
Go的内置集合中没有set 可以 map[type]bool
bool 类型的默认零值为false

1 set 元素唯一性
2 基本操作
添加元素
元素唯一性
删除元素
元素个数

python 中的集合类型
list 列表
dictory 字典
set 集合
tuple 元祖


*/

func TestMapWithFunc(t *testing.T)  {
	m := map[int]func(op int)int{}
	m[1] = func(op int)int{return op}
	m[2] = func(op int)int{return op*op}
	m[3] = func(op int)int{return op*op*op}
	t.Log(m[3](4))
	t.Log(m[2](5))
}

// 使用map 实现set数据类型

func TestMapForSet(t *testing.T)  {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is  not existing",n)
	}
	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet,1)
	t.Log(len(mySet))
	if mySet[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is  not existing",n)
	}

}