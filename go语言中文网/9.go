package main

import "fmt"

func main(){
	//for i:=0;i<=10;i++{
	//	if i > 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Printf("\nline after for loop")
	for i:=0;i<=10;i++{
		if i%2 == 0{
			continue
		}
		fmt.Println(i)
	}

}