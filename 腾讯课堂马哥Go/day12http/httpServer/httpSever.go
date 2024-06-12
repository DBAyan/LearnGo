package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func boyHandler(w http.ResponseWriter, r *http.Request)  {
	io.Copy(os.Stdout,r.Body)
	fmt.Fprintf(w,"hello boy!\n")
	for k,v := range r.Header {
		fmt.Printf("%v = %v\n", k,v)
	}
	for _ ,ck := range r.Cookies() {
		fmt.Printf("%s = %s\n", ck.Name,ck.Value)
	}
}

func girlHandler(w http.ResponseWriter, r *http.Request)  {
	io.Copy(os.Stdout,r.Body)
	fmt.Fprintf(w,"hello girl!\n")
	for k,v := range r.Header {
		fmt.Printf("%v = %v\n", k,v)
	}
}
func main()  {
	// 定义路由
	http.HandleFunc("/boy",boyHandler)
	http.HandleFunc("/girl",girlHandler)

	// 把http服务器启动
	http.ListenAndServe(":6868", nil)
}
