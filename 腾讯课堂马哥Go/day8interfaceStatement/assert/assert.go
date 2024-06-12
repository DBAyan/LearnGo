package main

import "fmt"

func assert (i interface{})  {
	switch v:= i.(type) {
	case int:
		fmt.Printf("type i is %T,value is %d\n", v,v)
	case float64:
		fmt.Printf("type i is %T,value is %f\n", v,v)
	case byte,string,rune:
		fmt.Printf("type i is %T,value is %v\n", v,v)
	}
}

func main()  {
	var i interface{}
	var a int = 3
	var b float64 = 3.14
	var c string = "中国人"

	i = a
	assert(i)

	i = b
	assert(i)

	i= c
	assert(i)
	
}
