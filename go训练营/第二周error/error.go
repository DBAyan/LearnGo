package main

import (
	"errors"
	"fmt"
)

/* 一 Error vs exception

Go error 就是一个普通的接口 返回一个普通的值

type error interface {
	Error() string
}

error 是个接口 ，接口就是方法的集合 ，里面有一个 名字叫 Error() string 的方法 。这个方法不接受任何参数，返回一个 string类型 。

errors.New 函数 接收一个字符串text，返回一个error接口类型的值（&errorString{text}），
errorString 是一个接口，有一个类型为string的s字段

errorString 有方法  Error() string ，实现了接口  error。

package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.

func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

错误  交给调用者处理

异常 fatal error 不可恢复的程序错误 数组的索引越界 ，栈错误，不可恢复的环境问题

 */


type Notifier interface {
	notify()
}

type WechatNotifier struct {
	Name string
	Message string
}

type QQNotifier struct {
	Name string
	Message string
}

type EmailNotifier struct {
	Name string
	Message string
}

func (w *WechatNotifier) notify ()  {
	fmt.Printf("%v notify %v",w.Name,w.Message)
}

func (q *QQNotifier) notify()  {
	fmt.Printf("%v notify %v",q.Name,q.Message)
}

func (e *EmailNotifier) notify()  {
	fmt.Printf("%v notify %v",e.Name,e.Message)
}

func sendNotify (notifier Notifier)  {
	notifier.notify()
}


// 布尔类型只能有两个状态 比如0是正数还是负数 ?

func Positive(n int) bool {
	return n > -1
}

func Check(n int)  {
	if Positive(n) {
		fmt.Println(n,"is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

// 利用 go语言的多返回值

func Positive2(n int)  (bool,bool){
	if n == 0 {
		return false, false
	}
	return n > -1 ,true
}

func Check2(n int)  {
	pos,ok := Positive2(n)
	if !ok {
		fmt.Println(n,"is neither")
		return
	}
	if pos {
		fmt.Println(n,"is positive")
	} else {
		fmt.Println(n,"is negative")
	}
}

// 使用error 简单的约定 + 多参数返回

func Positive3(n int) (bool,error){
	if n == 0 {
		return false,errors.New("undefined")
	}
	return n > -1,nil
}

func Check3(n int)  {
	pos , err := Positive3(n)
	if err != nil {
		fmt.Println(n,err)
		return
	}
	if pos {
		fmt.Println(n,"is positive!")
	} else {
		fmt.Println(n,"is negative!")
	}
}

// 单返回值 不建议这种方式 返回值具有二义性

func Positive4(n int) *bool {
	if n == 0 {
		return nil
	}
	r := n > -1
	return &r
}

func Check4(n int) {
	pos := Positive4(n)
	if pos == nil {
		fmt.Println(n,"is neither!")
		return
	}
	if *pos {
		fmt.Println(n,"is positive!")
	} else {
		fmt.Println(n,"is negative!")
	}
}

// 使用panic 与 recover 模拟 try catch 的写法

func Positive5 (n int) bool {
	if n == 0 {
		panic("undefined")
	}
	return n > -1
}

func Check5(n int)  {
	defer func() {
		if recover() != nil {
			fmt.Println(n,"is neither!")
		}
	}()
	if Positive5(n) {
		fmt.Println(n,"is positive!")
	} else {
		fmt.Println(n,"is negative!")
	}
}


func main ()  {
	w := &WechatNotifier{
		Name: "微信",
		Message: "微信消息",
	}

	q := &QQNotifier{
		Name: "QQ",
		Message:"QQ消息",
	}

	e := &EmailNotifier{
		Name: "Email",
		Message: "邮件消息\n",
	}

	sendNotify(w)
	sendNotify(q)
	sendNotify(e)

	Check(1)
	Check(0)
	Check(-1)

	Check2(-1)
	Check2(0)
	Check2(1)

	Check3(-1)
	Check3(0)
	Check3(1)

	Check4(-1)
	Check4(0)
	Check4(1)

	Check5(-1)
	Check5(0)
	Check5(1)
}