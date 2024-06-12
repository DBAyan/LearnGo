package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main()  {
	get()
	post()
	complexHttpRequest()
}

// get请求
func get()  {
	resp, err :=http.Get("http://127.0.0.1:8888/boy")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // 一定要调用 resp.Body.Close() ，否则会协程泄露
	io.Copy(os.Stdout,resp.Body)
	// 打印 响应头
	for k,v := range resp.Header {
		fmt.Println(k," = ", v)
	}
	fmt.Println(resp.Status) // 响应状态
	fmt.Println(resp.Proto) // http协议
}

// post请求
func post()  {
	reader:= strings.NewReader("hello server") // 新建一个io.Reader类型
	resp , err := http.Post("http://127.0.0.1:8888/girl","text/plain",reader) // 第一个参数是URL，第二个参数是 contentType 类型，第三个参数是请求正文，并不是字符串，而是io.Reader类型
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout,resp.Body)
	defer resp.Body.Close()
	// 打印resp.Header 响应头
	for k,v := range resp.Header {
		fmt.Println(k, "==>", v)
	}
}

// 复杂的请求
func complexHttpRequest() {
	reader := strings.NewReader("hello server")
	// 创建请求，该函数接受三个参数 分别为请求方法，请求的url ,body
	req , err := http.NewRequest("POST","http://127.0.0.1:8888",reader)
	if err != nil {
		panic(err)
	}
	// 自定义请求头
	req.Header.Add("User-Agent","中国")
	req.Header.Add("MyHeaderKey","MyHeaderValue")
	// 自定义cookie
	req.AddCookie(&http.Cookie{
		Name:"yhh",
		Value: "yhh_pwd",
		Path:"/",
		Domain: "localhost",
		Expires: time.Now().Add(time.Duration(time.Hour)),
	})

	// 构建client
	client := &http.Client{
		Timeout: 100 * time.Millisecond, // 设置请求的超时时间, 100毫秒 。
	}

	// 提交http请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// 一定要记得关闭
	defer resp.Body.Close()

	// 打印resp中的内容
	io.Copy(os.Stdout,resp.Body)

	// 打印resp header中的内容
	for k,v := range resp.Header {
		fmt.Println(k," = ", v)
	}


}