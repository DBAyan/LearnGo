package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

//  go get -u github.com/julienschmidt/httprouter


func handler(method string,w http.ResponseWriter,r *http.Request)  {
	fmt.Printf("your method is %s\n", method)
	fmt.Fprintf(w,"your method is %s\n", r.Method)
	fmt.Println("request body:")
	io.Copy(os.Stdout,r.Body)
	w.Write([]byte("hello Booy!"))

	for k,v  := range r.Header {
		fmt.Printf("%s = %s\n", k, v)
	}

	fmt.Println(r.Body)
}

func get (w http.ResponseWriter,r *http.Request, params httprouter.Params) {
	// 处理panic
	//str := "123"
	//fmt.Println(str[3])
	handler("get",w,r)

}

func post (w http.ResponseWriter,r *http.Request, params httprouter.Params) {
	handler("POST",w,r)

}

// RESTful 传递参数，把参数当做路径的一部分  :参数 *参数, * 号可以匹配多级路径
func handlerName(w http.ResponseWriter,r *http.Request, params httprouter.Params)  {

	fmt.Printf("Your name is %s,Your user type is %s,your addr is %s\n",params.ByName("name"),params.ByName("type"),params.ByName("addr"))
	fmt.Fprintf(w,"Your name is %s,Your user type is %s,your addr is %s\n",params.ByName("name"),params.ByName("type"),params.ByName("addr"))
}

func handlerPanic(w http.ResponseWriter,r *http.Request, i interface{})  {
	fmt.Fprintf(w,"Server Panic %v",i)
}

func main()  {
	router := httprouter.New()
	router.GET("/",get)
	router.POST("/",post)

	router.GET("/user/:name/:type/*addr",handlerName)

	// 返回静态文件
	// 当又请求访问类似 localhost:8888/file/a.html 的路径的时候，返回给用户 服务器路径  ./static  的a.html
	// 访问不存在的文件 比如  localhost:8888/file/b.html ,会报 404 page not found
	// ServerFiles 只能通过GET 方法请求

	router.ServeFiles("/file/*filepath",http.Dir("./static"))

	// 支持对panic的处理

	router.PanicHandler = handlerPanic

	http.ListenAndServe(":8888",router)
}
