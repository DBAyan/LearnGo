package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
//const (
//	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
//	O_RDONLY int = syscall.O_RDONLY // open the file read-only. 只读
//	O_WRONLY int = syscall.O_WRONLY // open the file write-only. 只写
//	O_RDWR   int = syscall.O_RDWR   // open the file read-write. 读写
//	// The remaining values may be or'ed in to control behavior.
//	O_APPEND int = syscall.O_APPEND // append data to the file when writing. 在文件中追加内容
//	O_CREATE int = syscall.O_CREAT  // create a new file if none exists. 创建文件
//	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.  和O_CREATE一起使用 ，表示在创建新文件时，如果文件已经存在，则打开文件失败，返回一个错误。如果不使用O_EXCL标志，当文件已存在时会打开该文件并且不会清空其内容。
//	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O. 指定在写入文件时进行同步操作。
//	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened. 打开文件时清空文件内容。
//)

func main()  {

	file := createFile("234.text")
	// 文件中写入数据 方式一 使用io.WriteString 函数
	if n,err := io.WriteString(file,"hello world\n");err != nil {
		fmt.Printf("写入文件失败 报错 %v",err)
	} else {
		fmt.Printf("写入文件成功 %v", n)
	}
	// 文件中写入数据 方式二 使用 结构体中的Write函数
	file.Write([]byte("hello yanhaihang\n中国人\n"))
	// 文件中写入数据 方式三  使用 bufio.NewWriter
	writer := bufio.NewWriter(file)
	writer.WriteString("hello gogogogo\n")
	writer.Flush()
	//removeFile("234.text")

	readFileioutil()
	readFileBufio()
}



// 打开文件用来只读的操作 。 os.Open 调用了方法 OpenFile(name, O_RDONLY, 0)
// file 是指针变量

func openFile()  {
	file,err := os.Open("/Users/yanhaihang/Desktop/go/LearnGo/learn_self/123.log")

	if err != nil {
		fmt.Printf("打开文件失败 报错 %v",err)
	} else {

		fmt.Println("打开文件成功")
		fmt.Println(file)
	}
}


// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
// 创建文件 如果使用goland的执行会创建的工作目录下 使用go run 会创建在当前目录下

func createFile(fileName string)  *os.File  {
	file,err := os.Create(fileName)
	if err != nil {
		fmt.Printf("创建文件失败 报错 %v",err)
	} else {
		fmt.Println("创建文件成功")
		fmt.Printf("%v",file)

	}
	return file
}

// 删除文件
func removeFile(fileName string)  {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("删除文件失败 报错 %v",err)
	} else {
		fmt.Println("删除文件成功")

	}
}


func readFileioutil()  {
	byteDate,err := ioutil.ReadFile("234.text")
	if err != nil {
		fmt.Printf("读取文件失败%v",err)
	} else {
		fmt.Println(byteDate)
		fmt.Println(string(byteDate))
	}
}

func readFileBufio()  {
	file,err := os.Open("234.text")
	if err != nil {
		fmt.Printf("读取文件失败%v",err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line,err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("读取失败")
				break
			}
			fmt.Println(line)
		}
	}
}