package main

import (
	"fmt"
)


func calBill(price int,no int) int{
	var totalPrice = price * no
	return totalPrice
}

func calProps(width float64,length float64) (float64,float64){
	var area = width * length
	var perimeter = (width + length) * 2
	return area,perimeter
}

func rectProps2(width float64,length float64)(area,perimeter float64){
	area = width * length
	perimeter = (width + length) * 2
	return
}
func main(){
	price,no := 25,4
	totalPrice := calBill(price,no)
	fmt.Println("The total price is $",totalPrice)
	width,length := 1.2, 2.4
	area,perimeter := calProps(width,length)
	fmt.Printf("Area %f,Perimeter %f\n",area,perimeter)

	area3,_ := calProps(1,7.8)
	fmt.Printf("Area %f\n",area3)

	area2,perimeter2 := rectProps2(5.5,6.7)
	fmt.Printf("Area %f,Perimeter %f",area2,perimeter2)

}