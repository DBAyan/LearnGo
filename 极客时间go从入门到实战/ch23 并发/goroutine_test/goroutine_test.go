package goroutine_test

import (
	"testing"
	"time"
)

/*
goroutine 协程
协程 VS 线程
携程的创建调度机制 

*/


func TestGoroutine(t *testing.T)  {
	for i := 0;i < 10;i++ {
		go func(i int) {
			println(i)
		}(i)
	}
	time.Sleep(100*time.Millisecond)
}