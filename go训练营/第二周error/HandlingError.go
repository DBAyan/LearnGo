package main

import (
	"fmt"
	"os"
)

func main()  {
	v,err := os.Open("/Users/yanhaihang/Desktop/go/LearnGo/go训练营/第二周error/testfil")
	if err != nil {
		fmt.Println(err) //  open  no such file or directory
	}
	fmt.Println(v) // &{0xc000054180}

}
