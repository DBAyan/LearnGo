package main

import (
	"fmt"
)

func main(){
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("===分割线===")

	t2 := i.(string)
	fmt.Println(t2)

}
