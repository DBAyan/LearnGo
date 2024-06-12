package main

import (
	"fmt"
	"time"
)

//
//func hello(){
//	fmt.Println("hello world")
//}
//
//func main(){
//	go hello()
//	time.Sleep(1)
//	fmt.Println("main func")
//
//}



func numbers(){
	for i:=1;i<=5;i++{
		time.Sleep(240 * time.Microsecond)
		fmt.Printf("%d",i)
	}
}

func alphabets(){
	for i:='a';i<='e';i++{
		time.Sleep(400*time.Microsecond)
		fmt.Printf("%c",i)
	}
}

func main()  {
	go numbers()
	go alphabets()
	time.Sleep(3000*time.Microsecond)
	fmt.Print("main terminated")
}