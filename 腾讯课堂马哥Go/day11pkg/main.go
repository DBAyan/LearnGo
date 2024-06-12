package main

import (
	mypkg "day11pkg/pkgDirctory01" // 有别名
	"day11pkg/pkgDirctory02"       // 没有别名
	_ "day11pkg/pkgDirctory03" //
	"fmt"
)


func init()  {
	fmt.Println("in pkg main DDD")
}

func main()  {
	type Student struct {
		id int
		info pkgDirctory02.User // pkg_name.User
	}

	var student01 = Student{1, pkgDirctory02.User{"yhh", 30, mypkg.Address{"山东", "郓城"}}}
	fmt.Println(student01)
}

func init()  {
	fmt.Println("in pkg main CCC")
}


