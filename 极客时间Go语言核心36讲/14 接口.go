package main

import "fmt"

/*
一 如何判断数据类型的方法 实现了就是 接口类型中的方法？
1 两个方法的签名(参数类型 与 返回值类型)一致
2 两个方法的名称形同

二
*/

type Pet interface {

	// SetName(name string)

	Name() string
	Category() string
}

type Dog struct {
	name string
}

// *Dog 指针类型拥有3个方法
// Dog 值类型拥有2个方法

func (d Dog)Name() string {
	return d.name
}

func (d *Dog)SetName(name string)  {
	d.name = name
}

func (d Dog)Category() string {
	return "dog"
}

func main()  {
	// 短变量声明一个变量 dog ,是Dog类型，并且进行了赋值初始化
	dog := Dog{"BaGong"}
	// &dog 叫做pet变量的 实际值（动态值），*Dog类型叫做pet变量的 实际类型（动态类型）
	// pet变量的 静态类型 是Pet
	//var pet Pet = &dog

	//_, ok := interface{}(dog).(Pet)
	//fmt.Printf("Dog implements interface Pet: %v\n", ok)

	fmt.Printf("The dog's name is %q.\n",dog.Name()) // The dog's name is "BaGong".
	var pet Pet = dog
	dog.SetName("DaHuang")
	fmt.Printf("The dog's name is %q.\n",dog.Name()) // The dog's name is "DaHuang".
	fmt.Printf("The pet is a %s. the name is %q\n",
		pet.Category(),pet.Name()) // The pet is a dog. the name is "BaGong"
	fmt.Println()

	// 实例2
	dog1 := Dog{
		"XiaoHei",
	}
	fmt.Printf("The name of first dog is %q.\n",dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "XiaoBai"
	fmt.Printf("The name of first dog is %q.\n",dog1.Name()) // The name of first dog is "XiaoBai".
	fmt.Printf("The name of second dog is %q.\n", dog2.Name()) // The name of second dog is "XiaoHei".
	fmt.Println()

	// 实例3
	dog = Dog{"DaHuang"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())

	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}


