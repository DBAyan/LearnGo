package main

import (
	"fmt"
	"time"
)

//func main(){
//	//var c chan int
//	c := make(chan int)
//	if c == nil{
//		fmt.Println("channel a is nil, going to define it")
//		//c = make(chan int)
//		fmt.Printf("Type  of c is %T",c)
//	} else {
//		fmt.Println(c)
//	}
//}

func hello(done chan bool){
	fmt.Println("hello go routine is going to sleep ")
	time.Sleep(4 * time.Second)
	fmt.Println("hello world goroutine")
	done <- true
}

func calcSquares(number int , squareop chan int){
	sum:=0
	for number != 0{
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubesop chan int){
	sum := 0
	for number != 0{
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubesop <- sum
}


func fornmber(number int){
	sum := 0
	for number !=0 {
		// 取余数 例如 123%10 = 3
		digit := number % 10
		fmt.Println(digit)
		sum += digit * digit
		fmt.Println(sum)
		// 取商 例如 123/10 = 12
		number /= 10
		fmt.Println("2",number)
	}
}
func main(){
	//done := make(chan bool)
	//fmt.Println("main function goint to call hello gp routine")
	//go hello(done)
	//<- done
	//fmt.Println("main function")
	//fmt.Println(123%10)
	//fornmber(123)
	number := 589
	sqrch := make(chan int)
	cubch := make(chan int)
	go calcSquares(number,sqrch)
	go calcCubes(number,cubch)
	squares, cubes := <- sqrch,<-cubch
	fmt.Println("Final output", squares + cubes)

}


