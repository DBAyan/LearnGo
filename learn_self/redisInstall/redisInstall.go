package main

import (
	"flag"
	"fmt"
)

var (
	redisVersion string = ""
)
func main()  {
	flag.StringVar(&redisVersion,"redisVersion", "", "安装Redis的版本")
	//flag.StringVar()
	osOpt()
}


// 操作系统参数优化
func osOpt()  {
	command := []string{"sh","-c", "echo never > /sys/kernel/mm/transparent_hugepage/enabled"}
	fmt.Println(command[0:])

	fmt.Printf("%T",command[1:])
	for idx,value := range command[1:] {
		fmt.Printf("索引 %d 的值为 %s\n", idx, value)
	}
}