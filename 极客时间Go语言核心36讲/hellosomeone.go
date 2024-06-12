package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var address string

func init(){
	// 接受参数 1 接受的变量 2 变量名字 3 默认值 4 参数解释
	flag.StringVar(&name,"name","everyone","输入某人的名字")
}

func main()  {
	// --help 自定义输出
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "欢迎某人的脚本\n")
		flag.PrintDefaults()
	}

	//真正的解析参数
	flag.Parse()
	fmt.Printf("hello %s! \n",name)
}

