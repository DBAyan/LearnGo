package main

import "fmt"

//在 Go 语言中，iota 并没有一个正式的缩写含义，它只是一个被预定义的标识符，用于在常量声明中生成递增的常量值序列。
//根据 Go 语言官方文档的描述，iota 是一个希腊语单词，表示为“第九个字母（I）”。在数学中，希腊字母的顺序通常被用来表示一系列连续的值，因此 iota 也被用来表示连续的常量值序列。
//虽然 iota 并没有一个官方的缩写含义，但是在 Go 社区中，一些开发者经常将其解释为“枚举器（enumerator）”，或者“索引器（indexer）”等。这些解释都与 iota 在常量声明中生成连续的序列值有关。



//在 Go 语言中，iota 是一个被预定义的标识符，它被用来在常量声明中生成一系列相关值。

//具体来说，iota 在常量声明中用于表示递增的常量值序列。在常量声明中，每当遇到一个新的常量关键字 const 时，iota 的值都会被重置为 0，然后每出现一次 iota，它的值就会递增一次。




const (
	FATAL int = iota // 0
	CRITICAL // 1
	ERROR //2
	NOTICE //3
	INFO //4
	DEBUG //5
)

const (
	x = 1<<iota // 左移0位 1
	y // 左移1位 2
	z // 左移2位 4
	o // 左移3位  8
)

func main()  {
	fmt.Println(FATAL)
	fmt.Println(CRITICAL)
	fmt.Println(ERROR)
	fmt.Println(DEBUG)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(o)
}


