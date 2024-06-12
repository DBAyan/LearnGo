package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name string
	salary int
	currenry string
}

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

func (e Employee) diaplaySalary() {
	fmt.Printf("salary of %s is %s%d",e.name,e.currenry,e.salary)
}

func (r Rectangle) Area() int {
	return r.length*r.width
}

func (c Circle) Area() float64{
	return c.radius*c.radius*math.Pi
}

func main() {
	emp1 := Employee{
		name:"yanhaihang",
		salary:5000,
		currenry: "$",

	}
	r:= Rectangle{
		length: 5,
		width: 10,
	}
	c:=Circle{
		radius: 4.6,
	}

	emp1.diaplaySalary()
	fmt.Printf("Area of rectangle %d\n", r.Area())
	fmt.Printf("Area of circle %f", c.Area())
}
