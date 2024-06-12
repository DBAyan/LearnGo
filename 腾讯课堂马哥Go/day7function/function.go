package main

import "fmt"

// 一 形参与实参
// a b 是形参 ，形参数函数内部的局部变量，实参的值会拷贝给形参
// 形参可以有0个或者多个
// 参数类型相同时可以只写一次   例如 add(a,b int)

func add(a,b int) int  { // 注意 大括号 不能另起一行
	a = a + b  // 在函数内部修改形参的值 ，实参不会受到影响
	return a
}

// 二 传递指针类型参数
// 如果想通过函数修改实参的值，需要传递指针类型
func argf(a,b *int) int {
	*a = *a + *b
	*b = 888
	return *b
}

// 三 传引用和传引用的指针
// slice map channel三大引用类型，slice其实是一个含有三个字段的struct。
// 他们作为函数参数时，其实跟普通struct没有什么区别，都是对struct内部各个字段做一次拷贝传到函数内部

func sliceArgs(arr []int)  { // slice 作为参数 ，其实就是（想象成传三个参数）把arrayPointer(int指针)，len(int),cap（int）拷贝一份传进来
	arr[0] = 111 // 修改底层数组的首元素，因为把底层数组的指针传递进来了，所以修改底层数组 是修改的对原始的内存空间产生影响
	arr = append(arr,2) // arr的len,cap发生变化，不会影响实参 ,append在函数内部是值传递
	fmt.Printf("arr : %v \nlength of arr is %d,\ncapacity of arr is %d\n",arr, len(arr),cap(arr))

}

//  四 函数返回值
//可以有0个或者多个返回值
//可以在func行直接声明要返回的变量,但是必须加括号
//return后面的语句不会执行
//无返回值时return可以不写

func returnFunc (a,b int) (c int)  { // 返回变量c已经声明好了
	c = a +b
	return // 这里不需要再写c了
}

// 五 不定长参数
// 不定长参数 append 就是传的不定长参数
// 不定长参数可以不传   variablesEngthArg(9)
func variablesEngthArg(a int, other ...int) (sum int) { // 把不定长参数当做slice 类型
	sum = a
	for _, ele := range other{
		sum += ele
	}
	fmt.Printf("len of other is %d, cap of other is %d\n", len(other), cap(other)) // 可以使用slice的一些函数 比如len,cap等
	return
}

// Byte(int8 类型别名)类型的切片，调用时的形式 variByteArg(arrByte...)。
// byte加三个点 对应byte类型的切片

func variByteArg(other ...byte) (sum byte) {
	sum = 0
	for _,eles := range other{
		sum += eles
	}
	fmt.Println(sum)
	return
}

// 六 递归函数 斐波那锲数列
//形式：
//0 1 1 2 3 5 8 13 21 34
// 自己实现(循环法) https://www.jianshu.com/p/a71924dedbc3
func fibo(n int) (fb []int) {
	fb = append(fb,0)
	for i,a,b:=1, 0, 1; i < n; i++ {
		fb = append(fb,b)
		a,b = b ,a + b
	}
	fmt.Printf("fb:%v\n",fb)
	return
}

// 利用递归实现 F(0)=0 F(1)=1 F(n)=F(n-1)+f(n-2)
func F(n int) int {
	if n == 0|| n==1 {
		return n
	}
	return F(n-1) + F(n-2)
}

// 七 匿名函数
// 函数作为参数（变量）使用
func foo(a,b int)  int {
	return a+b
}
func funArg(f func(a,b int) int, b int) int { // f 参数是一种函数类型 ，零值为nil
	a := 2*b
	return f(a,b)
}

// 可以简写如下 更加简介
type bar func(a,b int) int

func funArg2(f bar,b int) int {
	a := 2*b
	return f(a,b)
}

type  User struct {
	Name string
	age bar // age d的类型是 bar,而bar代表一种函数类型
	hello func(name string) string // 使用匿名函数l来声明
}

var sum3 = func(a,b int) int { return a + b }
//type add func(a,b int) int
//
//func op(f add, a int) int {
//	b:=2*a
//	return f(a,b)
//}



// 八 闭包函数
// 函数作为返回值
// sub 内的局部便令 i
func sub() func() {
	i:= 10
	fmt.Printf("i pointer:%p\n",&i)
	b:= func() {
		fmt.Printf("i addr:%p\n",&i)
		i--
		fmt.Println(i)
	}
	return b
}


func main()  {
	x,y := 2,3 // x，y 是实参
	c := add(x,y)
	fmt.Println(c)

	m,n := 666,777
	argf(&m, &n)
	fmt.Println(m,n)

	//var arr []int
	//arr = make([]int,2,3)

	//arr := make([]int,2,3)

	//arr := []int{0,2}
	//fmt.Printf("arr : %v \nlength of arr is %d,\ncapacity of arr is %d\n",arr, len(arr),cap(arr))
	//
	//sliceArgs(arr)
	//fmt.Printf("arr : %v \nlength of arr is %d,\ncapacity of arr is %d\n",arr, len(arr),cap(arr))

	brr := make([]int,2,3)
	fmt.Printf("brr : %v \nlength of brr is %d,\ncapacity of brr is %d\n",brr, len(brr),cap(brr))

	sliceArgs(brr)
	fmt.Printf("brr : %v \nlength of brr is %d,\ncapacity of brr is %d\n",brr, len(brr),cap(brr))

	r := returnFunc(1,2)
	fmt.Println(r)

	sum := variablesEngthArg(1, 2,3,4)
	fmt.Println(sum)
	sum2 := variablesEngthArg(9)
	fmt.Println(sum2)

	arrByte := []byte{2,4,6}
	variByteArg(arrByte...)

	s := "abc"
	variByteArg([]byte(s)...)

	slice1 := append([]byte("Hello"), "wolrd"...) // "wolrd"... 自动把"world"转成byte切片，等价于[]byte("world")...
	fmt.Println(slice1)
	slice2 := append([]rune("abc"),[]rune("ABC")...)  // 需要显示的把"ABC" 转换为rune切片
	fmt.Println(slice2)

	// 斐波那锲数列
	fibo(1)
	fibo(10)

	fmt.Println(F(10))

	// 调用函数作为参数的函数

	f := funArg(foo,3)
	fmt.Println(f)


	ch := make(chan func( string) string,10)
	ch <- func(name string) string {
		return "hello" + name
	}

	sum3(1,2)

	//闭包
	ff := sub()
	ff()
	ff()

}