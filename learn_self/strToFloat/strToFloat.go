package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换为小数
	str := "3.14"
	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
		fmt.Println("转换出错:", err)
		return
	}

	// 打印转换后的小数
	fmt.Printf("转换后的小数: %f\n", f)
}

