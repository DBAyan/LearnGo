package main

import (
	"fmt"
	"time"
)

// 定义结构体类型

type User struct {
	id int
	name,addr string // 多个字段类型相同时可以简写为一行
	score float64
	enrollment time.Time
}

// 匿名结构体 只能创造一个变量
var people struct{
	Name string
	sex int8
}

type UserMap map[int]User

func (um UserMap ) getid(no int) int {
	return um[no].id
}

// 结构体方法
func ( user User) hello(man string) string  {
	return "hello " + man +", i am " + user.name
}

// 匿名方法 不需要用到结构体中的字段 下划线写不写都可以
func (_ User) think(man string) string  {
	return man + "is thinking!"
}

func hello2(man string,user User) string  {
	return  "hello " + man +", i am " + user.name
}

func main()  {

	// 声明 与 初始化
	var u User // 声明变量
	u = User{id:1, name: "yhh", addr: "beijing",score: 99.5} // 只赋部分值

	u1 := User{} // 初始化为字段相应的零值
	u2 := User{2,"nico","xinjiang",65.4,time.Now()} // 省略字段，但是赋值时的顺序必须严格一致
	fmt.Println(u)
	fmt.Println(u1)
	fmt.Println(u2)

	u.enrollment = time.Now()
	u.score = 60
	fmt.Println(u.enrollment)

	fmt.Println(u.hello("nico"))
	fmt.Println(u.think("yhh"))
	fmt.Println(hello2("yhh",u))

	people.Name="xiaojiao"
	fmt.Println(people)



}


