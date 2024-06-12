package share_mem_test

import (
	"sync"
	"testing"
	"time"
)
/* 线程安全
sync.Mutex
Rwlock 读写锁分开了

sync.WaitGroup
含义 ：
只有当我wait的所有东西都完成之后 ，那么我才能继续往下执行

原来的都是通过sleep 是担心外面执行的协程速度超过了里面协程的速度

*/


func TestCounter(t *testing.T)  {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter ++
		}()
	}
	time.Sleep(100 * time.Millisecond)
	t.Logf("counter is %d", counter)
	// counter is 4836
	// why?? 因为采用的counter在不同的协程中自增 ，导致并发的竞争条件，不是线程安全的程序，
	//要安全的会就需要对共享内存进行锁的保护
	// 非线程安全丢失了很多正确的写操作


}

func TestCounterTreadSafe(t *testing.T)  {
	counter :=0
	var mut sync.Mutex

	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}() // defer 延迟最后执行 释放锁
			mut.Lock() // 加锁
			counter ++ // +1
		}()
	}
	time.Sleep(100 * time.Millisecond)
	t.Logf("counter is %d", counter) // counter is 5000
}

func TestCounterWaitGroup(t *testing.T)  {
	counter :=0
	var mut sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1) // 增加要等待的数量
		go func() {
			defer func() {
				mut.Unlock()
			}() // defer 延迟最后执行 释放锁
			mut.Lock() // 加锁
			counter ++ // +1
			wg.Done() // 去告诉waitgroup这个完成了
		}()
	}
	wg.Wait() //  2.这里等待上面所有的等待完成后马上走下面的程序 , 这样就可以准确的等待协程执行了
	t.Logf("counter is %d", counter) // counter is 5000
}