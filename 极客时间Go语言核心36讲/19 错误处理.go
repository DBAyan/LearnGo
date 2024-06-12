package main

import (
	"errors"
	"fmt"
)
/*

一  普通错误处理
无法预测的错误处理 panic  recover
error 包中的 New函数  输入字符串 返回接口 地址&errorString{text}
func New(text string) error {
	return &errorString{text}
}
*/

// 具名返回值 return 省略了，返回值 err 是接口类型
func echo (request string) (response string, err error) {
	// go编程思想 错误尽早处理返回
	if request == "" {
		err = errors.New("empty request")
		fmt.Printf("Type of err : %T", err) //  *errors.errorStringerror: empty request
		return
	}
	response = fmt.Sprintf("echo %s",request)
	return 
}

func main()  {
	for _, req := range []string{"","hello world"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n",err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}