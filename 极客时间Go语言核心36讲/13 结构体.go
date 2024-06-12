package main

import "fmt"

// 动物分类学中基本分类法
// 定义结构体类型

type AnimalCategory struct {
	kingdom string
	phylum string
	class string
	order string
	family string
	genus string
	species string
}

type Animal struct {
	scientificName string
	AnimalCategory // 嵌入字段 匿名字段 ，只有类型 没有字段的名称
}

// AnimalCategory类型的 String方法

// 【默认】在 Go 语言中，我们可以通过为一个类型编写名为String的方法，来自定义该类型的字符串表示形式。
//这个String方法不需要任何参数声明，但需要有一个string类型的结果声明。

func (ac AnimalCategory)String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s-%s-%s",
		ac.kingdom,ac.phylum,ac.class,ac.order,ac.family,ac.genus,ac.species)
}

func (a Animal)String() string {
	return fmt.Sprintf("%s (category:%s)",a.scientificName,a.AnimalCategory)
}

func main()  {
	// 短变量声明 初始化

	//动物界、脊索动物门、哺乳纲、食肉目、猫科、豹属、虎种
	//Animal kingdom, chordate, Mammalia, Carnivora, Felidae, leopard and tiger species
	category := AnimalCategory{kingdom: "Animal",phylum: "chordate",class: "Mammalia",order: "Carnivora",
		family: "Felidae",genus: "leopard",species: "tiger"}

	// 在调用fmt.Printf函数时，使用占位符%s和category值本身就可以打印出后者的字符串表示形式，
	//而无需显式地调用它的String方法。
	fmt.Printf("The animal catagory :%s\n", category)

	// 只要名称相同，无论这两个方法的签名是否一致，被嵌入类型的方法都会“屏蔽”掉嵌入字段的同名方法。
	animal := Animal{
		scientificName: "Northeast Tiger",
		AnimalCategory: category,
	}
	// 虽然 Animal 类型没有方法 String方法，但是调用了 AnimalCategory 类型的String方法
	fmt.Printf("The animal: %s\n", animal)
}