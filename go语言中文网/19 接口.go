package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	 empId int
	 basicpay int
	 pf int
}

type Contract struct {
	empId int
	basicpay int
}

func(p Permanent) CalculateSalary() int{
	return p.basicpay+p.pf
}

func(c Contract) CalculateSalary() int{
	return c.basicpay
}

func tatalExpnds(s []SalaryCalculator){
	expense := 0
	for _,v := range s{
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per mouth is $%d",expense)
}

type Test interface {
	Tester()
}
type myFloat float64

func (m myFloat) Tester() {
	//fmt.Printf("Interface type %T value %v\n",t,t)
	fmt.Println(m)
}

func describe(t Test){
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func describe2(i interface{}){
	fmt.Printf("Type = %T, value = %v\n", i, i)
}


func main(){
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	empployees := []SalaryCalculator{pemp1, pemp2, cemp1}
	tatalExpnds(empployees)

	fmt.Println("")
	var t Test
	f := myFloat(78.9)
	t = f
	describe(t)
	t.Tester()

	// 空接口
	s := "hello world"
	describe2(s)

	i := 55555
	describe2(i)

	strt := struct {
		name string
	}{name: "yanhaihang"}
	describe2(strt)

}