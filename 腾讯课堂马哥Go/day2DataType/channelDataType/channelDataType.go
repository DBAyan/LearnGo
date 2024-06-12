package main

import "fmt"

func main()  {
	var  ch chan int // 声明
	fmt.Printf("ch is nil :%t\n",ch == nil)

	ch = make(chan int, 8) // 初始化 可容纳8个元素
	fmt.Printf("ch length is %d\n",len(ch))
	fmt.Printf("ch capa is %d\n",cap(ch))

	// 向管道中添加元素
	ch <- 3

	fmt.Printf("ch  is %v\n",ch)
	fmt.Printf("ch length is %d\n",len(ch))
	fmt.Printf("ch capa is %d\n",cap(ch))

	for i:=0;i<7;i++{
		ch <-i
	}
	fmt.Printf("ch  is %v\n",ch)
	fmt.Printf("ch length is %d\n",len(ch))
	fmt.Printf("ch capa is %d\n",cap(ch))

	// 从管道中取走元素
	v := <- ch
	fmt.Println(v)
	v1 := <- ch
	fmt.Println(v1)

	// 遍历管道中的元素 遍历的同时也会取走元素

	close(ch) // for range 遍历之前一定要关闭管道

	for ele := range ch{
		fmt.Println(ele)
	}

	// 等价于 这样取走不用close
	fmt.Printf("ch length is %d \n",len(ch))
	L := len(ch)
	for j:=1; j<=L; j++ {
		ele:= <- ch
		fmt.Println(ele)
	}

	fmt.Printf("ch length is %d \n",len(ch))

}


