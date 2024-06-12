package operator_test

import (
	"fmt"
	"testing"
)

/*
一
go语言没有前置++ --  ++a --a
go语言只有后置++ -- a++

二
go语言中 数组的比较规则 ==

相同维度且含有相同个数的数组才可以比较
每个元素都相同的才想等

三
逻辑运算符
逻辑与
逻辑或
逻辑非

四
位运算符
按位清零运算符
&^

*/
func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3}
	b := [...]int{1,3,4}
	e := [...]int{1,3,2}
	//c := [...]int{1,2,3,4}
	d := [...]int{1,2,3}
	t.Log(a==b) // false 长度（元素个数）相同，元素不同，两个数组不相等
	t.Log(a==e) // false  元素的顺序不同，两个数组不相等

	// 维度不同的数组不能比较 ，如果比较会有编译错误
	//t.Log(a==c)  invalid operation: a == c (mismatched types [3]int and [4]int)
	t.Log(a==d) // true  长度 元素顺序 完全相同 ，数组才相等
}

const (
	Readable = 1 << iota // 0001 1 iota 0
	Writeable // 0010 2 iota 1
	Executable //  0100  4 iota 2
)

func TestBitClear(t *testing.T){
	fmt.Println(Readable,Writeable,Executable)
	a := 7 //0111
	a = a &^ Readable
	t.Log(a&Readable == Readable,a&Writeable == Writeable,a&Executable == Executable)

}