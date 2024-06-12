package main

import "fmt"

func matrix_mu()  {
	
}

func breakFor() {
	arr := []int{11,22,33,44,55}
	for i,ele := range arr {
		fmt.Println(i,ele)
		if i > 2 {
			//break
			continue
		}
		fmt.Printf("next:%d\n", i)
	}
}

func main()  {
	arr := []int{11,22,33,44,55}
	for i:=0; i<len(arr) ;i++ { // for 声明初始化局部变量; 条件表达式; 后续操作
		fmt.Printf("%d:%d\n",i,arr[i]) // for 局部变量仅在for块内可见
	}
	//i = i+1 for块之外不可见

	//i := 0 // 只有条件判断时 前后的;可以省略
	//for i<len(arr)  {
	//	fmt.Printf("%d:%d\n",i,arr[i])
	//	i ++ // 后续操作可以移动到for语句的{}
	//}

	//for { // for {} 无限循环
	//	fmt.Println("无限循环")
	//}

	// 比较复杂的初始化 或者 后续操作有多个表达式
	for sum,j := 0,0;j< len(arr) && sum < 100;sum,j = sum + arr[j],j+1 {
		fmt.Printf("sum:%d，i:%d\n",sum,j)
	}

	// 遍历数组 遍历切片
	for idx,ele := range arr{
		fmt.Printf("arr[%d] is %d\n",idx,ele)
	}
	// 遍历字符串 string
	str := "中国人aAbBcC"
	for idx , ele:= range str{ // ele 是rune类型
		fmt.Printf("%d:%d\n",idx,ele)
	}
	//遍历map go不保证顺序
	booklist := make(map[int]string, 10)
	booklist[1] = "《高性能MySQL》"
	booklist[2]= "《Redis设计与实现》"
	booklist[3]= "《go》"
	booklist[4]= "《python》"
	for k,v := range booklist {
		fmt.Printf("序号%d:%s\n",k,v)
	}
	//遍历channel 遍历之前一定要close
	var c chan string
	c = make(chan string,8)
	c <- "aa"
	c <- "bb"
	c <- "cc"
	close(c)
	for ele := range c{
		fmt.Printf("%s\n",ele)
	}

	breakFor()
}
