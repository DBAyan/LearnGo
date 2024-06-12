package main

/*
c语言的错误处理
1 显式处理每个错误 代码更健壮
2
不足 ：
一值多用

Go
函数多返回值

// fmt包
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

error  接口类型表示错误 ，如果没有错误 则返回nil .如果有错误 则返回具体的错误。
源码

type error interface {
	Error() string
}

*/
