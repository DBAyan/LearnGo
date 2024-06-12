package loop_test

import "testing"

/*
循环三段式
for i=1;i++;i<5{}

使用for 关键字实现 while 循环


 */

func TestLoop(t *testing.T){
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}