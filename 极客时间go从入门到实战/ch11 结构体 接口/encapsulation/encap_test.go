package encap

import (
	"fmt"
	"testing"
	"unsafe"
)
/*
go 语言是面向对象的编程语言吗？
官方 Yes or No
DUCK TYPE


面向对象编程
结构体 定义
注意的是 每个field后不需要逗号
type struct_name struct {
Id int
Name string

}
实例创建以及初始化


封装 对于数据与行为的封装


一 成员（值）定义
二 行为（方法）定义

无论 receiver 是传值还是指针 都可以正常使用该方法

区别

传指针 是同一块内存
 Address of e is c00006e490
 Address of e is c00006e490

传值 会在内存中复制一块新的
 Address of e is c00006e3a0
 Address of e is c00006e3d0

建议使用传指针

使用指针类型作为 receiver没有内存复制 节约内存


*/
type Employee struct {
	Id string
	Name string
	Age int
}

func (e Employee)StringFormat() string {
	fmt.Printf(" Address of e is %x\n",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}


//func (e *Employee)StringFormat() string {
//	fmt.Printf(" Address of e is %x\n",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
//}

func TestCreateEmployeeObj(t *testing.T)  {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Miko",Age: 30}
	e2 := new(Employee)
	e2.Id = "1"
	e2.Name = "yanhaihang"
	e2.Age = 30
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2) // &{1 yanhaihang 30}
	t.Logf("e type is %T",e) // e type is encap.Employee encap 包名 Employee 结构体名
	t.Logf("e2 type is %T",e2) // e2 type is *encap.Employee 指针类型

}


func TestStructOperations(t *testing.T)  {
	e := Employee{"0", "Bob", 20}
	fmt.Printf(" Address of e is %x\n",unsafe.Pointer(&e.Name))
	t.Log(e.StringFormat())
}