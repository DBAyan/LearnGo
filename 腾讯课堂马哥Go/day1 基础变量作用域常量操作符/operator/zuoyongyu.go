package zuoyongyu1

import "fmt"

//全局变量必须通过var  const 定义，不能通过:= 定义

var (
	AA = 11 // 大写开头 全局变量
)

const (
	Pi = 3.14
	E = 2.17
)

func Scope2()  {
	fmt.Println(BB)
}
