package csp_test

import (
	"fmt"
	"testing"
	"time"
)

/*
CSP并发控制方式
70年代左右提出
communicating sequential processes ：依赖通道完成两个协程之间协调


ACTor model 两个通讯实体之间发送信息的机制来协调
区别

ACTor model ：直接通信的机制 不需要channel
CSP 需要 channel

Channel基本机制
通讯的两方必须都在通道上 ，接收消息不在 ，放消息的人也会被阻塞。接收消息同理，如果放消息的不在
buffered Channel

松耦合
放消息的人就放 如果容量满了，必须拿出一个消息
接收消息同理 

*/

func service() string  {
	time.Sleep(time.Millisecond * 100)
	return "DONE"
}

func otherTask()  {
	fmt.Println("working on someting else!")
	time.Sleep(time.Millisecond*50)
	fmt.Println("Task is Done")
}

func TestService(t *testing.T)  {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string{
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second*1)
}