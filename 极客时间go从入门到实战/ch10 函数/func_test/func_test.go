package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
自己需要补充的知识点 ：
格式换的符号

go语言中的== := =的区别
= 是赋值语句
:=是声明类型（编译器帮你做的）并赋值

== 是判断两个值是否相等

函数 ：一份公民
一 函数特性
1 函数可以有多个返回值
传值还是传引用？？

2 所有参数都是值传递
slice map channel 会有传引用的错觉

3 函数可以作为变量的值,函数可以作为map的值
4 函数可以作为参数或者返回值，可以很大长度上简化我们的程序

计算函数的运算时长

二 变长参数
将变长参数转换为数组（切片的底层数据接口也是数组）
func funcName(elems ...int) {}

三 延迟执行函数 defer
什么时候执行defer后的函数？
调用了某个函数 不会立即执行 ，先执行下面的函数 ，等return之前再执行defer的函数
有什么用呢 ：
逻辑处理完成后 关闭资源 释放锁 清理资源等 避免程序锁住 资源浪费等

panic 不可修复的错误
即使有panic ,defer也是会执行的
但是在panic之后的逻辑不会执行

说到了匿名函数 ？
*/

func returnMultiValues()(int,int){
	return rand.Intn(100),rand.Intn(10)
}

// 计算函数执行的时长
// 参数是个函数类型 返回也是函数类型
// function program 类似于Python中的装饰器

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		fmt.Println(start)
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T){
	a,_:=returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(22))
}

// 变长参数 ...type

func sum(elems ...int) int {
	ret := 0
	fmt.Printf("%T\n", elems) // []int slice 切片
	fmt.Println(elems)
	for _,elem := range elems {
		fmt.Println(elem)
		ret += elem
	}
	return ret
}

func TestVarParam(t *testing.T)  {
	t.Log(sum(1,2,3))
	t.Log(sum(55,44,66))
}

func Clear()  {
	fmt.Println("Clear resource!")
}

func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("start!")
	panic("err") // 即使有panic  defer 函数也是会执行的。
	fmt.Println("end!")
}