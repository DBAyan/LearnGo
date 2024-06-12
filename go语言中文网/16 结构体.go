package main

import "fmt"
//import "structs"

func main()  {
	type Employee struct {
		firstname string
		lastname string
		age int
	}

	type Student struct {
		firstname,lastname,grade string
		age int

	}

	var employee2 struct{
		firstname,lastname string
		age,salary int
	}

	fmt.Println(Employee{})
	fmt.Println(employee2)

	//Employee 's type is main.Employee
	//,employee2's type is structs { firstname string; lastname string; age int; salary int }

	fmt.Printf("Employee 's type is %T\n,employee2's type is %T\n",Employee{},employee2)

	student1 := Student{
		firstname : "haihang",
		lastname:"yan",
		age : 30,
		grade : "junior",
	}
	student2 := Student{"jia","peng","junior",29}

	student3 := struct {
		firstname,lastname,grade string
		age int
	}{
		firstname: "san",
		lastname: "zhang",
		grade: "senior",
		age: 15,
	}

	fmt.Printf("info of student1 is %v\n",student1)
	fmt.Println(student2)
	fmt.Println("student3 is" ,student3)

	fmt.Printf("san zhang's info\n" +
		"firstname of zhangsan is %v\n" +
		"lastname of zhangsan is %v\n",student3.firstname,student3.lastname)

	// 匿名结构体，只有类型没有字段名，
	type Person struct {
		string
		int
	}
	p := Person{"yanhaihang",35000}
	fmt.Println(p)
	// 字段名默认为字段的类型
	fmt.Println(p.string)

	// 嵌套结构体
	type Address struct {
		city,state string
	}
	type Person2 struct {
		name string
		age int
		addr Address
	}
	var p2 Person2
	p2.name = "yanhaihang"
	p2.age = 30
	p2.addr = Address{"heze","yuncheng"}
	fmt.Println("name:",p2.name)
	fmt.Println("age:",p2.age)
	fmt.Println("city:",p2.addr.city)
	fmt.Println("state:",p2.addr.state)

	//提升字段 如果结构体中有匿名的结构体字段，则该匿名结构体里的字段就称为提升字段，
	//这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问。
	type PersonInfo struct {
		habbit,job string
	}
	type Person3 struct {
		name string
		age int
		PersonInfo
	}
	p3 := Person3{
		"yhh",30,PersonInfo{
			"game","DBA",
		},
	}

	fmt.Println(p3)
	fmt.Println("name:",p3.name)
	fmt.Println("age:",p3.age)
	// 可以直接访问
	fmt.Println("job:",p3.job)

	// 可导出结构体和字段 首字母大写
	//var spec structs.Spec
	//spec.Maker = "apple"
	//spec.Price = "5000"
	//fmt.Println(spec)

	// 结构体比较
	type name struct {
		firstname string
		lastname string
	}
	name1 := name{"hh","y"}
	name2 := name{"hh","y"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are equal")
	}
	name3 := name{"hhh","y"}
	fmt.Println(name1==name3)
}

