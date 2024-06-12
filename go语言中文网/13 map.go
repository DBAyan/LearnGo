package main

import "fmt"

/*
可以理解为Python中的字典 ，键值对的数据结构。

一 map的初始化
1
make(map[type of key]type of value)

2 只声明
var personSalary map[string]int

3 声明并初始化

二 map的操作
增 删 查 改


 */

//persons lary := make(map)

func main()  {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil.Go to makeone.")
		fmt.Println(personSalary)
	}

	m1 := make(map[string]int)
	fmt.Println(m1)

	// 声明并初始化
	m2 := map[string]int{
		"zhangsan":3000,
		"lisi":4000,
		"yhh":35000,
	}
	fmt.Println(m2)

	// 添加元素
	m2["haihang"] = 50000
	fmt.Println(m2,len(m2))

	// 获取map中的元素
	fmt.Println("salary of haihang is ",m2["haihang"])
	// 如果该key不存在，则返回零值
	fmt.Println("salary of nico is ",m2["nico"])

	// 如何判断map中key是不存在还是这个key的value为0呢？
	value ,ok := m2["nico"]
	if ok{
		fmt.Println("salary of nico is", value)
	} else {
		fmt.Printf("empolyee %v is not exist!", "nico")
	}
}