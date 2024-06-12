package main


// 一 if 单分支 二分支 多分支

// 二 逻辑运算符 与 && 或 || 非 ! 以及优先级 ，建议使用带有小括号的子布尔表达式来清晰地表达判断条件。

// 三 支持声明 if 语句的自用变量
// 什么是自用变量 ；自用变量的形式与位置；

// 四 if的快乐路径原则
// 1 仅使用单分支控制结构；
// 2 当布尔表达式求值为 false 时，也就是出现错误时，在单分支中快速返回；
// 3 正常逻辑在代码布局上始终“靠左”，这样读者可以从上到下一眼看到该函数正常逻辑的全貌；
// 4 函数执行到最后一行代表一种成功状态。

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println(runtime.GOOS) // darwin
	fmt.Println(runtime.GOARCH) // amd64
	fmt.Println(runtime.Compiler) // gc

	if runtime.GOOS == "darwin" {
		fmt.Println("we are on darwin os")
	}
	if (runtime.GOOS == "darwin") && (runtime.GOARCH == "amd64") && (runtime.Compiler=="gc") {
		fmt.Println("we are using standard go compiler on darwin os for amd64")
	}

	a, b := false,true
	if a && b != true {
		fmt.Println("(a && b) != true")
		return
	}
	println("a && (b != true) == fales")
}
