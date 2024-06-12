package main

import (
	"fmt"
	"net/http"
)

func handlerHello(w http.ResponseWriter,r *http.Request)  { // 两个参数 ，将返回参数写入到 w, 请求参数在参数r中，这里是简单的例子，所有没有使用到r参数
	fmt.Println(r.Header)
	fmt.Println(r.Proto)
	fmt.Println(r.ProtoMajor)
	fmt.Println(r.ProtoMinor)
	fmt.Println(r.Form)
	fmt.Println(r.ContentLength)
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println(r.Body)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Host)
	fmt.Println(r.Response)
	fmt.Println(r.PostForm)
	defer r.Body.Close()
	// 打印Header中的内容
	for k,v := range r.Header {
		fmt.Println(k, "==>", v)
	}
	// 打印cookie的内容

	for _,cookie := range r.Cookies() {
		fmt.Printf("%s = %s", cookie.Name,cookie.Value)
	}
	fmt.Fprintf(w,"Hello Worald!") // 把返回内容写入 http.ResponseWriter

}

func handlerBoy(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello Boy")
}

func handlerGirl(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello girl")
}

func main()  {
	// 定义路由，将访问不同目录的请求 路由到 不同的处理函数
	http.HandleFunc("/",handlerHello) // 路由 ，访问 / 根目录是去执行 handlerHello,上面定义好的函数
	http.HandleFunc("/boy",handlerBoy) // 路由 ，访问/boy目录是去执行 handlerBoy
	http.HandleFunc("/girl",handlerGirl) // 第一个参数是个字符串 ，第二个参数是个函数

	// 启动HTTP server 服务，ListenAndServe 如果不发生error会一直阻塞。为每一个请求创建一个协程去处理
	if err := http.ListenAndServe(":8888",nil); err != nil { // 服务端口为 8888
		fmt.Printf("start http server fail : %s", err)
	}

}



