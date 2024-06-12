package err_test

import (
	"errors"
	"testing"
)

/*
try catch finally

go 语言支持多返回值
C 程序在返回结构的时候返回错误值
Go 的错误机制
1 没有异常机制
2 error类型实现了error接口
3 可以通过errors.New()
error.New("")
*/

// 斐波那锲数列

// 1 1 2 3 5 8

func GetFibocci(n int) ([]int, error)  {
	if n < 0 || n > 100 {
		return nil, errors.New("n must bi 0-100")
	}
	fibList := []int{1,1}
	for i := 2;i<n;i++{
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}

func TestFibo(t *testing.T)  {
	//ret,err := GetFibocci(10)
	//t.Log(ret)
}