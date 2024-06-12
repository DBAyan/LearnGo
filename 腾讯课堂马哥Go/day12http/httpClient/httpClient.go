package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)


// get 方法

func get()  {
	resp,err := http.Get("http://localhost:8888/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)

	// 状态以及话术
	fmt.Println(resp.Status) // !200 OK

	// 打印 response.Header
	//Date=[Tue, 13 Dec 2022 08:06:53 GMT]
	//Content-Length=[11]
	//Content-Type=[text/plain; charset=utf-8]
	for key,value := range resp.Header {
		fmt.Printf("%v=%v\n", key,value)
	}
	// 协议
	fmt.Println(resp.Proto) // HTTP/1.1

}

func post()  {
	reader := strings.NewReader("Hello Server!")
	resp , err := http.Post("http://localhost:8888/","text/plain", reader)
	if err != nil {
		fmt.Println("发送请求失败", err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(resp.Proto)
	for k , v := range resp.Header {
		fmt.Printf("%v : %v\n",k,v)
	}
}

// 复杂的HTTP 请求
func complexHttpRequest()  {
	reader := strings.NewReader("hello http server !")
	req , err := http.NewRequest("POST","http://localhost:6868/boy",reader)
	if err != nil{
		panic(err)
	} else {
		// 自定义请求头
		req.Header.Add("User-Agent","中国")
		req.Header.Add("MyHeaderKey","MyHeaderValue")
		// 自定义Cookie
		req.AddCookie(&http.Cookie{
			Name: "auth",
			Value: "passwd",
			Path: "/",
			Domain: "localhost",
			Expires: time.Now().Add(time.Duration(time.Hour * 1)),
		})
		client := &http.Client{
			Timeout: 100 * time.Millisecond,
		}
		if resp , err := client.Do(req); err != nil { // 提交http请求
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			io.Copy(os.Stdout,resp.Body)
			fmt.Println(resp.Status)
			fmt.Println(resp.Proto)
			for k , v := range resp.Header {
				fmt.Printf("%v : %v\n",k,v)
			}
		}

	}

}

func main()  {
	get()
	//post()

	//complexHttpRequest()
}
