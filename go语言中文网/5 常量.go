package main

import "fmt"

func main(){
	const a = 55
	//a = 89  cannot assign to a (declared const)
	var name = "sam"
	fmt.Printf("type %T value %v",name,name)
}
