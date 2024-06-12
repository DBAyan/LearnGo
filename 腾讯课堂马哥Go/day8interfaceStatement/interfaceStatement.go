package main

import "fmt"

// 接口的定义
// 自己认为  接口是一类事物公有的方法 ，比如 动物都可以叫 ，人类可以说话，交通工具都有一个速度

//type Transpoter interface { // 接口一般以er结尾
//	move(string, string)(int)
//}
//
//type Car struct {
//	color string
//	speed int
//	monoy float64
//}
//
//func (car Car) move(src,dest string) int  {
//	fmt.Printf(" The car can move from %s to %s\n", src,dest)
//	return 10
//}
//
//type Bike struct {
//	color string
//	monoy float64
//}
//
//func (bike Bike) move(src,dest string) int  {
//	fmt.Printf(" The car can move from %s to %s\n", src,dest)
//	return 30
//}

type Animals interface {
	sayHello()(string)
}
type Dog struct {
	name string
}

func (dog Dog) sayHello() string {
	return fmt.Sprintf("%s say hello is 汪汪", dog.name)
}

type Cat struct {
	name string
}

func (cat Cat) sayHello() string {
	return fmt.Sprintf("%s say hello is 喵喵", cat.name)
}

type Human struct {
	name string
}

func (human Human) sayHello()  string {
	return fmt.Sprintf("%s say hello is 你好", human.name)
}

func sayHello(a Animals)  string {
	fmt.Println(a.sayHello())
	return  a.sayHello()
}

func main()  {

	//var car Car
	//car = Car{"white",200,2000000}
	//bike := Bike{"yellow", 800}
	//car.move("霍营","西二旗")
	//bike.move("霍营","西二旗")
	//
	//var transpoter Transpoter
	//transpoter = car
	//transpoter.move("霍营","西二旗")

	var jinmao Animals
	jinmao = Dog{"金毛"}
	fmt.Println(jinmao.sayHello())

	var huban Animals = Cat{"虎斑"}
	fmt.Println(huban.sayHello())

	var chinese Animals = Human{"中国人"}
	fmt.Println(chinese.sayHello())

	sayHello(jinmao)
	sayHello(huban)
	sayHello(chinese)


}
