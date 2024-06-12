package main

import (
	"fmt"
	"structtest/bookdir"
	"sync"
	"unsafe"
)


// 声明
// 赋值
// 修改
// 查询


func main()  {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
	type T1 int
	type T2 T1
	type T3 string
	var n1 T1
	var n2 T2 = 5
	n1 = T1(n2)
	fmt.Println(n1)
	//var s T3 = "hello"
	//n1 = T1(s) // cannot convert s (type T3) to type T1
	var b bookdir.Books
	b.Title = "The Go Programming Language"
	b.Pages = 800

	type Empty struct{}

	var s Empty
	fmt.Println("the length of s is ",unsafe.Sizeof(s)) //  0
	//var c = make(chan Empty)
	//c <- Empty{}

	type Person struct {
		Name string
		Phone string
		Addr string
	}

	type article struct {
		Title string
		Author Person
	}
	type article2 struct {
		Title string
		Person
	}

	var a article
	a.Title = "jikeshijian"
	a.Author.Phone = "13501223450"
	phone := a.Author.Phone

	fmt.Println("The phone number of author is ",phone)

	var a2 article2
	fmt.Println("The phone number of author is ",a2.Phone)
	//var boo boo
	//fmt.Println(boo)
	type book struct {
		name string
		pages int
		indexes map[string]int
	}

	var b2 = book {
		"MySQL实战45讲",
		45,
		map[string]int{"lock":3},
	}


	fmt.Println(b2) // {MySQL实战45讲 45 map[lock:3]}
	fmt.Println(b2.indexes)
	b2.name = "Redis核心技术与实战"
	fmt.Println(b2.name) // Redis核心技术与实战

	type info struct{
		name string
		habbit string
		salary float32
		family map[string]string
	}

	i := info{name: "yhh",habbit: "basketball",salary: 357200.32,family: map[string]string{"mom":"myt","father":"yts"}}
	fmt.Println(i)

}
