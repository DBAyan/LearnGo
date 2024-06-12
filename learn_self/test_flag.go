package main

// 在Go中，flag是一个标准库，用于解析命令行参数。它提供了一种简单和方便的方式来接受和处理命令行参数。下面是一些flag的详细用法：


import (
	"flag"
	"fmt"
	"os"
)

var (
	testValue string
	// 基本使用
	// 第一个参数是 flag 名称，第二个参数是默认值，第三个参数是 flag 的说明。flag.Parse() 函数会解析命令行参数，并将其存储在对应的变量中。
	// 布尔类型的命令行参数
	version = flag.Bool("version",false, "string参数")
	help = flag.Bool("help",false,"帮助信息")
	// 字符串
	host = flag.String("host","127.0.0.1","服务器IP")

	// Port 整数
	Port = flag.Int("port",3306,"数据库端口")
	execFlag = flag.Bool("flag",true,"二进制参数")




)

func main() {

	flag.Parse()

	// StringVar defines a string flag with specified name, default value, and usage string.  命令行中参数的名字，默认值， 用法
	// The argument p points to a string variable in which to store the value of the flag. 第一个参数是一个指针变量（取普通变量的地址），用来存储命令行参数的值
	flag.StringVar(&testValue,"testStr","aaa","测试StringVar")
	fmt.Println(testValue)

	// 版本信息 和 帮助信息 在程序的逻辑中，
	//我们检查 version 和 help 变量是否被设置为 true，如果是，则打印相应的信息并退出程序。否则，程序将继续执行其余逻辑。
	if *version{
		fmt.Println("my program version :1.0")
		return}

	//if *help {
	//	fmt.Println("Usage:myprogram [OPTION]")
	//	fmt.Println("  --version Print version infomation and exit")
	//	fmt.Println("  --help print USage information and exit")
	//	return
	//}

	// flag.PrintDefaults 打印所有的命令行参数 类型 解释 默认值
	if *help {
		fmt.Fprintf(os.Stdout,"Usage of gh-ost:\n")
		flag.PrintDefaults()
	}


	fmt.Println(*host)
	fmt.Println(*execFlag)
	fmt.Println(*Port)


}
