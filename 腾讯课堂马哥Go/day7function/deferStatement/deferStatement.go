package main

import "fmt"

//defer 用于注册一个延迟调用 （在函数返回之前调用）
//defer 典型的应用场景是释放资源 比如释放文件的句柄 关闭数据库连接
//如果同一个函数有多个defer,则后注册的先执行
//defer后可以跟一个func，func内部发生panic，main协程不会exit,其他defer可以正常执行
//defer后不是跟func,而是直接跟一条执行语句 ，则相关变量在执行defer时拷贝或计算


// a , d, c,b
func basic()  {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("d")
}


func deferFunc() (i int) {
	i = 9
	defer func() { // defer 可以跟func，在函数内部不会把9赋值 给 i,在返回的时候才进行赋值
		fmt.Printf("defer1 i :%d\n", i) // 5
	}()

	defer func(i int) {
		fmt.Printf("defer2 i :%d\n", i) //  9
	}(i) // i 作为实参传递进来

	defer fmt.Printf("defer3 i :%d\n", i) // 9  后面跟语句 ，在注册defer的时候，就会把9赋值给i（变量在注册defer的时候进行赋值或者拷贝 ）
	return 5 // 把5赋值为i
}

func deferPanic()  {
	n := 0
	fmt.Println("OO")
	defer fmt.Println("AA") //
	defer func() {
		fmt.Println(1 / n) // 除0异常 panic: runtime error: integer divide by zero
		fmt.Println("zzz")
		defer func() { // 不会被执行
			fmt.Println("CC")
		}()
	}()

	defer func() {
		fmt.Println("DD")
	}()
}


func main()  {
	basic()

	deferFunc()

	deferPanic()
}
