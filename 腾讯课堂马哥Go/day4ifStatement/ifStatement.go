package main

import (
	"fmt"
	"time"
)

func main()  {
	// if 语句
	if 9 > 5 {
		fmt.Printf("9>5 is %t\n", 9 > 5 )
	}

	if a,b,c := 2,5,9; b<c && (b>3 || b<a) {
		fmt.Println("fit")
	}

	// if - else 语句
	color := "green"

	if color == "read" { // if 只能有一个
		fmt.Println("color is Red!")
	} else if color == "blue" { // else if可以有0个或者多个
		fmt.Println("color is blue!")
	} else if color == "yellow" {
		fmt.Println("color is blue!")
	} else {  // else 可以有0个或者1个
		fmt.Println("I don`t know the color is")
	}

	// 通过if 判断map的key 是否存在

	booklist := make(map[int]string, 10)
	booklist[1] = "《高性能MySQL》"
	booklist[2]= "《Redis设计与实现》"
	for i:=1; i<=10;i++ {
		if value,exist := booklist[i]; exist {
			fmt.Printf("no %d is book %s\n", i,value)
		} else {
			fmt.Printf("no %d is not exist\n", i)
		}
	}

	// 多层嵌套 判断公交车道
	now := time.Now()
	hour := now.Hour()
	weekday := now.Weekday()
	if weekday.String() != "Starday" && weekday.String() != "Sunday" {
		if hour >= 7 && hour <=9 {
			fmt.Println("公交车道不让走")
		} else {
			fmt.Println("司家车可以走走")
		}
	} else {
		fmt.Println("司家车可以走走")
	}


}


