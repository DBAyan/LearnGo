package main

import (
	"fmt"
	"unicode/utf8"
)
func main(){
	// go 原生支持字符串的优点一 string 类型的数据是不可变的，提高了字符串的并发安全性和存储利用率。
	var a string = "hello world"
	//a[0] = 'k' // cannot assign to a[0] (strings are immutable)
	a = "gopher"
	fmt.Printf(a)

	// go 原生支持字符串的优点二 `` 反引号 所见即所得
	var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(s)

	var s2 = "中国人"
	fmt.Printf("the length of s2 is %d\n",len(s2)) // 9 字节数 byte

	for i := 0;i <= len(s2);i++ {
		fmt.Printf("0x%x ", s[i])
		fmt.Printf("\n")
	}

	// 3 字符数
	fmt.Println("The character count in s2 is ",utf8.RuneCountInString(s2))

	for _,c := range s2{
		fmt.Printf("0x%x ", c)
	}

	// 下标操作
	fmt.Printf("0x%x\n", s2[0]) //  0xe4

	// rune 类型
	// type rune = int32
	var c rune = '严'
	fmt.Println(c)  // 输出内容为unicode的码点（位次，唯一值）。输出的是十进制 20005，转换为十六进制为 4e25。\u4e25 可以转换为字符
	var d rune = 'a'
	fmt.Println(d) // 97--> 0061 \u0061 --> a

	// 单引号''中的字面值  输出的都是十进制的Unicode的码点
	fmt.Println('a')// 97
	fmt.Println('严') // 20005
	fmt.Println('\n') // 10
	fmt.Println('\'') // 39

	// \u 2个十六进制 \U 4个十六进制 ，输出的也是十进制的Unicode码点
	fmt.Println('\u4e25')  // 20005
	fmt.Println('\U00004e2d') // 20013
	fmt.Println('\u0027') // 39

}