package constant_test

import "testing"

// 连续常量赋值

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)


func TestConstantTry(t *testing.T){
	t.Log(Monday,Tuesday,Wednesday)
}