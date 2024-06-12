package main

import "fmt"

type  QuackableAnimal interface {
	Quack()
}

type Duck struct {

}

func (d Duck) Quack()  {
	fmt.Println("duck quack!")
}

type Dog struct {

}
func (d Dog) Quack()  {
	fmt.Println("dog quack!")
}

type Bird struct {

}

func (b Bird) Quack ()  {
	fmt.Println("Bird quack!")
}

func AnimalQuackInForest(a QuackableAnimal)  {
	a.Quack()
}

func main()  {
	animals := []QuackableAnimal{new(Duck),new(Dog),new(Bird)}
	for k,animal := range animals{
		fmt.Println(k)
		AnimalQuackInForest(animal)
	}
}