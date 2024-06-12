package extension

import (
	"fmt"
	"testing"
)

/*
匿名嵌套类型

*/
// Pet 父类
type Pet struct {
	
}

func (p *Pet) Speak()  {
	fmt.Println("Wang!")
}

func (p *Pet) Speakto(host string)  {
	p.Speak()
	fmt.Println("",host)
}

type Dog struct {
	//p *Pet
	Pet // 匿名字段
}


//func (d *Dog) Speak()  {
//	d.p.Speak()
//}
//
//func (d *Dog) Speakto(host string)  {
//	d.p.Speakto(host)
//}

func TestDog(t *testing.T)  {
	dog := new(Dog)
	dog.Speakto("yhh")
}
