package map_test

/*
一 map的初始化
方式一
初始化并赋值
m1 := map[int]int{1:1,2:4,3:9}

方式二
初始化
m2 := map[int]int{}

方式三
make关键字

map不能使用capacity访问

二 map元素的访问
在访问key不存在时，仍会返回零值，

不能通过返回nil来判断元素是否存在
可以通过map多返回值if判断。

v,ok := m[key] ok 是个布尔值 ，如果key存在则返回true ,不存在则返回false。

三 map的遍历
for k,v:= range m {
}
返回两个返回值 第一个是key  第二个是value

四 Map与工厂模式

函数是一等公民


*/
import (
	"testing"
)

func TestMapInit(t *testing.T){
	m := map[string]int{"one":1,"two":2,"three":3}
	t.Log(m)
	t.Log("length of m:", len(m))

	m1 := map[int]int{1:1,2:4,3:9}
	t.Log(m1)
	t.Logf("length of m1 is %d",len(m1))

	m2 := map[int]int{}
	m2[5] = 25
	t.Log(m2)
	t.Logf("length of m2 is %d",len(m2))

	m3 := make(map[int]int,10)
	t.Log(m3,len(m3))
}


func TestAccessNotExistingKey(t *testing.T){
	m1 := map[int]int{} // 空map
	v,ok := m1[1] // 不存在的key返回该类型的零值 这里就是 整型的零值 0
	t.Log(m1[1]) // 0
	t.Log(v,ok) //  0 false

	m1[2] = 0 // key为2的 value值为0
	v1,ok1 := m1[2]
	t.Log(v1,ok1) // 0 true
	t.Log(m1[2]) // 0

	if v,ok := m1[4];ok{
		t.Logf(" %v key 2 is %d",m1,v)

	}else {
		t.Logf("key 4 is not exist!")
	}
}

// map 遍历 for range
func TestMapTravel(t *testing.T){
	m1 := map[string]int{"one":1,"two":22,"three":333}
	for k,v := range m1{
		t.Log(k,v)
	}
}
