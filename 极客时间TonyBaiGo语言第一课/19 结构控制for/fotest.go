package main

/*
一 for的经典形式
for 循环前置语句（仅执行一次）;条件判断表达式;循环后置语句 {
	循环体
}

二 go
*/
import "fmt"

func main()  {

	// 经典格式
	for a:=0; a<=10; a++ {
		fmt.Println(a)
	}

	// Go 语言的 for 循环支持 声明多循环变量，并且可以应用在循环体以及判断条件中
	for sum, i, j ,k := 0, 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k =i+1, j+1, k+5 {
		sum += (i + j + k)
		fmt.Println(sum)
	}

	// 省略循环后置语句
	for i := 0; i<10; {
		fmt.Println(i)
		i++
	}

	// 省略循环前置语句
	x := 0
	for ; x < 10; x++ {
		fmt.Println("x:",x)
	}

	// 省略循环前置与后置语句
	y := 0
	for ; y < 10; {
		fmt.Println("y:", y)
		y++
	}

	// 都省略 无限循环

	//for {
	//	//
	//	//}
	//	//
	//	//for true {
	//	//
	//	//}

	// for range 格式 语法糖
	// 切片 slice
	sl := []int{11,22,33,44,55}
	for i:=0; i<len(sl);i++ {
		fmt.Printf("sl[%d] = %d\n",i,sl[i])
	}

	for i,v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}

	for i := range sl {
		fmt.Println(i)
	}

	for _,v := range sl {
		fmt.Println(v)
	}

	for range sl {

	}

	// string 字符串
	var s string = "中国人"
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n",i,string(v),v)
	}

	// map
	var m = map[string]int {
		"yhh": 30,
		"xiaojuan": 22,
		"nico": 29,
	}

	for k,v := range m {
		fmt.Printf(" %s : %d\n",k,v)
	}

	// continue

	var sum int
	sl2 := []int{11,22,33,44,55}

	for i := 0;i<len(sl2); i++ {
		if sl2[i] % 2 == 0 {
			continue
		}
		sum += sl2[i]
	}
	fmt.Println(sum)

	var sum2 int
	sl3 := []int{111,222,333,444,555}

	loop:
		for i := 0;i<len(sl3); i++ {
			if sl3[i] % 2 == 0 {
				continue loop
			}
			sum2 += sl3[i]
		}
	fmt.Println(sum2)

	// break

}
