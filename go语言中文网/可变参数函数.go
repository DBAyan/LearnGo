package main

import "fmt"
// 可变长参数 ，数据类型为切片 ，nums ...int

func find(num int,nums ...int){
	fmt.Printf("Type of nums if %T\n",nums)
	found := false
	for i,v := range nums{
		if v == num{
			fmt.Println(num,"found at index",i," in",nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num,"not found in ",nums)
	}
	fmt.Printf("\n")
}
func main(){
	find(1,2,3,1)
	find(55,66,44,77,55,99)
}

