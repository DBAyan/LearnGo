package main

import "fmt"

// ANSI颜色转义码
const (
	Black   = "\u001b[30m"
	Red     = "\u001b[31m"
	Green   = "\u001b[32m"
	Yellow  = "\u001b[33m"
	Blue    = "\u001b[34m"
	Magenta = "\u001b[35m"
	Cyan    = "\u001b[36m"
	White   = "\u001b[37m"
	Reset   = "\u001b[0m"
)

func main() {
	// 输出不同颜色的文本
	fmt.Println(Red + "This is red text" + Reset)
	fmt.Println(Green + "This is green text" + Reset)
	fmt.Println(Blue + "This is blue text" + Reset)

	if "0" == "0" {
		fmt.Println("test")
	}
}


