package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	countGuard sync.Mutex
)

func GetCount() int {
	countGuard.Lock()
	defer countGuard.Unlock()
	fmt.Println(count)
	return count
}

func SetCount(c int)  {
	countGuard.Lock()
	time.Sleep(10*time.Second)
	count = c
	countGuard.Unlock()
}

//func main()  {
//	go SetCount(1)
//	time.Sleep(1*time.Second)
//	go GetCount()
//	time.Sleep(15*time.Second)
//}


var mutex sync.Mutex				//互斥锁
func printer(str string){
	mutex.Lock()  //加锁,对什么加锁？？
	fmt.Println(mutex)

	defer func() {
		mutex.Unlock()
		fmt.Println(mutex)
	}()		//解锁

	for _,ch:=range str{
		fmt.Printf("%c",ch)
		time.Sleep(time.Millisecond*300)
	}
}
func user1(){
	printer("hello ")
}
func user2(){
	printer("world")
}
func main() {
	go user1()
	go user2()
	for  {
		;
	}
}