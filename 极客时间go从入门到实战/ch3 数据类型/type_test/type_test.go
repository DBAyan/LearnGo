package type_test

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

/*  go语言类型
1 byte = 8 bit 1字节 = 8位


bool 布尔 false  true
string 字符串
int int8 int16 int32 int64 整型 后面的数字代表机器的位数 如果使用int 在64位的机器上就是int64 在32位的机器上就是int32
https://blog.csdn.net/a6661314/article/details/122798788

uint uint8 uint16 uint32 uint64  uintptr 无符号整型
byte 字节 因为是8位 也是uint8的别名
rune 和字符串相关 是Unicode的编码值 uint32的别名
float32  float64 浮点
complex64 complex128 复数

类型转换
go语言对隐式类型转换非常严苛的

1 不支持隐式类型转换
2 别名和原有类型而不能进行隐式类型转换

类型的预定义值,
比如整型的最大值 浮点型的最大值 ，math包里预定义了这些值 可以直接用
math.MaxInt64

二 指针类型
go语言支持 垃圾回收机制 ，也支持 使用指针访问内存空间
但是go语言的指针使用是有一些限制的
1 不支持指针运算
2

*/


type MyInt int64

func TestImplicit(t *testing.T){
	var a int = 1
	fmt.Printf("%T\n",a)
	fmt.Println(unsafe.Sizeof(a)) // 8 Byte

	var b int64
	//b = a //报错 ：cannot use a (type int) as type int64 in assignment

	// 如何解决 强制类型转换 使用显示类型转换
	//b = int64(a)
	t.Log(a,b)
	var c MyInt
	//c = b //  cannot use b (type int64) as type MyInt in assignment
	c = MyInt(b)

	t.Log(a,b,c)


	// 类型的预定义值 int64的最大值
	t.Log(math.MaxInt64) // 9223372036854775807
	t.Log(math.MaxFloat32)

}

// 指针类型不能进行运算 go语言不支持指针运算

func TestPoint(t *testing.T){
	a := 1 // 整型变量a
	aPtr := &a // & 取址符 a的地址

	//aPtr = aPtr + 1 //invalid operation: aPtr + 1 (mismatched types *int and int)

	t.Log(a,aPtr) // 1 0xc000016288
	t.Logf("%T,%T",a,aPtr) //  int,*int
}

// 字符串的零值不是nil 为""

func TestString(t *testing.T){
	var s string // 只声明 但是没有初始化 默认值为""
	t.Log("*"+s+"*")
	t.Log(len(s))
	if s == ""{
		t.Log("字符串为空")
	}

}