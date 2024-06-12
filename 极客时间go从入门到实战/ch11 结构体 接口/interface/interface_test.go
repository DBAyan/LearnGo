package interface_test

import "testing"
/*
duck type
方法签名是相同的 就认为他们是相同

与其他语言的差异：
1 接口为非入侵的 实现不依赖与接口定义
2 所以接口的定义也可以包含在接口使用者的包内

接口变量
多态 接口

自定义类型

*/

// 声明一个interface
type Programmer interface {
	WriteHelloWorld() string
}

// 声明一个结构体
type GoProgrammer struct {

}

// 声明一个方法
func (g *GoProgrammer)WriteHelloWorld() string {
	return "fmt.Printfln(\"Hello World\")"
}

//
func TestInterface(t *testing.T)  {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

