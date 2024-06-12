package main

import (
	"flag"
	"fmt"
)

func main(){
	// 第一种方式
	//var name string
	//flag.StringVar(&name,"name","everyone","the greeting person")

	// 第二种方式
	//var name=flag.String("name","everyone","the greeting person")
	// 第三种方式
	name := flag.String("name","everyone", "The greeting object.")

	flag.Parse()
	fmt.Println(len(flag.Args()))
	fmt.Println(flag.Arg(0))
	//fmt.Printf("hello %v",name)
	// 第二种方式
	fmt.Printf("hello %v\n",*name)
}