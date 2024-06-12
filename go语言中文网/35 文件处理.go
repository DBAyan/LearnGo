package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

/*
// 编译
 go build "35 文件处理.go"

// 执行
./35\ 文件处理 -FilePath='/Users/yanhaihang/Desktop/go/LearnGo/go语言中文网/17 方法.go'

*/
func main(){
	fptr := flag.String("FilePath","/Users/yanhaihang/Desktop/go/LearnGo/go语言中文网/test.txt","The path of file ")
	// fmt.Println(fptr) //0xc000010270
	flag.Parse()
	fmt.Printf("The value of file path ",*fptr) // /Users/yanhaihang/Desktop/go/LearnGo/go语言中文网/test.txt

	data,err := ioutil.ReadFile(*fptr)
	if err != nil{
		fmt.Printf("File reading error", err)
		return
	}
	// fmt.Println("Contents of file:", data) // Contents of file: [104 101 108 108 111 32 103 111 32 119 111 114 108 100 33]

	fmt.Println("Contents of file:", string(data)) //  hello go world!

}
