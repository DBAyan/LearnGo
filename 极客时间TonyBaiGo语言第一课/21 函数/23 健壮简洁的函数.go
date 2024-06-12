package main

import (
	"fmt"
)

/*
一 健壮
1 不要相信任何外部的参数 -- 要做参数合法性校验
2 不要忽略任何一个错误 -- 要做错误处理
3 不要假定异常不会发生 -- go的异常处理机制
二 go 异常处理机制 panic
定义：
panic 指的是 Go 程序在运行时出现的一个异常情况。如果异常出现了，但没有被捕获并恢复，
来源
1 开发人员通过panic函数主动触发
2 go运行时
panicing

当bar 函数中调用了panic ，调用bar函数的foo 函数也相当于了调用了panic 后续的逻辑都会终止
call mian
call foo
call bar
panic: panic occur in bar

recover
使用条件 ： 必须放在defer 函数才能生效

call mian
call foo
call bar
recover the panic: panic occur in bar
exit foo
exit main

三 如何应对panic
1 评估程序对panic的忍受度
比如 go 脚本 与 http服务对panic的忍受度完全不同
局部不要影响整体

2 提示潜在的bug
C语言有 断言 assert
go语言没有assert 可以利用panic模拟assert 提示bug

在 Go 标准库中，大多数 panic 的使用都是充当类似断言的作用的。

3 不要混淆异常 与 错误
区分 error 与 panic
Python 中 try: exception: 相当于 error


四 简洁性
defer ：go语言中一种延迟调用机制 收尾工作

什么时候执行 ：
无论是执行到函数体尾部返回，还是在某个错误处理分支显式 return，又或是出现 panic，
已经存储到 deferred 函数栈中的函数，都会被调度执行。

目的：确保函数在异常或者错误退出之前 释放资源 释放锁

限制 ：defer 离不开函数 或者 方法
1 只有在函数或者方法内部才能使用defer
2 defer 关键字之后只能跟函数或者方法 deferred 函数
多个defer函数的执行顺序 先进后出

注意点
1 那些函数可以作为deferred函数
append、cap、len、make、new、imag 不能作为deferred 函数 可以包裹匿名函数来支持

defer func(){
	_ = append(sl,11)
}()

close、copy、delete、print、recover 可以作为deferred函数

2 注意 defer 关键字后面表达式的求值时机

3 defer 函数会带来性能开销
1.3版本之前开销性能较高 之后进行了性能优化


*/

func foo()  {
	println("call foo")
	bar()
	println("exit foo")
}

func bar()  {
	defer func() {
		if e := recover();e!=nil {
			fmt.Println("recover the panic:", e)
		}
	}()
	println("call bar")
	panic("panic occur in bar")
	zoo()
	println("exit bar")
}

func zoo()  {
	println("call zoo")
	println("exit zoo")
}

// 伪代码
//func dosometing() error {
//	var mu sync.Mutex
//	mu.Lock()
//
//	r1, err := OpenResource1()
//	if err != nil {
//		mu.Unlock()
//		return err
//	}
//
//	r2, err := OpenResourc2()
//	if err != nil {
//		r1.Close()
//		mu.Unlock()
//		return err
//	}
//
//	r3, err := OpenResourc3()
//	if err != nil {
//		r2.Close()
//		r1.Close()
//
//		mu.Unlock()
//		return err
//	}
//
//	err = doWithResource()
//	if err != nil {
//		r3.Close()
//		r2.Close()
//		r1.Close()
//		mu.Unlock()
//		return err
//	}
//
//	r3.Close()
//	r2.Close()
//	r1.Close()
//	mu.Unlock()
//	return nil
//
//}

func bar1() (int,int) {
	return 1, 2
}

func foo1()  {
	var c chan int
	var sl []int
	var m = make(map[string]int,10)
	m["item1"] = 1
	m["item2"] = 2
	//var a = complex(1.0,-1.4)

	var sl1 []int
	defer bar()
	//   defer discards result of
	//defer append(sl,11)
	//defer cap(sl)
	//defer imag(a)
	//defer len(sl)
	//defer make([]int,10)
	//defer new(*int)
	//defer complex(2,-2)
	//defer real(a)

	defer close(c)
	defer copy(sl1,sl)
	defer delete(m,"item2")
	defer print("hello,defer\n")
	defer println("hello defer")
	defer panic(1)
	defer recover()

}

func foo2()  {
	for i := 0; i <= 3; i++ {
		defer println(i)
	}
}

func foo3()  {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			println(n)
		}(i)
	}
}

func foo4()  {
	for i := 0; i <= 3; i++ {
		defer func() {
			println(i)
		}()
	}

}

func main()  {
	//println("call mian")
	//foo()
	//println("exit main")
	//
	//foo1()

	fmt.Println("foo2 result:")
	foo2()
	fmt.Println("foo3 result:")
	foo3()
	fmt.Println("foo4 result:")
	foo4()
}