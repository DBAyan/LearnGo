package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//标准输入输出
// 多个变量需要用空格分开
func stdFmt()  {
	fmt.Println("Please input your name and your class：")
	var s string
	var i int
	fmt.Scanf("%s\n", &s) // 这里为什么要有取地址符
	fmt.Scanf("%d\n", &i)
	fmt.Printf("Welcome %s to class %d!\n", s,i)

}

// 文件IO : 以字节的方式读取文件内容
func readFileByte()  {
	if file,err := os.Open("test.txt"); err != nil {  //  只写文件名的话需要与go文件在一个目录下 否则就写绝对路径，
		fmt.Printf("打开文件失败 %s\n", err)
	} else {
		defer file.Close()
		var buffer strings.Builder
		fmt.Println("打开文件成功")
		for {
			bs := make([]byte, 30) // 定义一个byte类型切片 len = cap = 10
			if n, err := file.Read(bs) ;err != nil {
				fmt.Printf("读取文件失败%s\n",err)
				if err == io.EOF {
					fmt.Printf("读取文件结束%s\n", err)
					break
				}
			} else {
				fmt.Printf("成功读取%d个字节\n", n)
				buffer.WriteString(string(bs))
			}
		}

		fmt.Println(buffer.String())
	}
}

//文件IO ：以字符串形式读取文件
func readFileString()  {
	if file,err := os.Open("test.txt"); err != nil {
		fmt.Printf("打开文件失败:%s", err)
	} else {
		defer file.Close()
		fmt.Println("open file success!")
		reader := bufio.NewReader(file)
		for {
			if line ,err := reader.ReadString('\n'); err!= nil{
				fmt.Printf("read file fail err:%s\n",err)
				if err == io.EOF { // End Of File 文件末尾
					fmt.Println("end of file ")
					break
				}
			} else {
				fmt.Println(strings.Trim(line,"\n"))
			}
		}

	}
}

//字节写文件
// os.O_CREATE 创建文件
// os.O_TRUNC 清空文件
//os.O_WRONLY 只能写不能读
// 取 或 | 就是
func writeFileByte() {
	if file,err := os.OpenFile("la01.text",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666); err != nil { // 也是打开文件 但是参数更多 控制更多
		fmt.Printf("open file fail ,err: %s", err)

	} else {
		defer file.Close()
		fmt.Println("open file success!")
		file.Write([]byte("hello go \n"))
	}
}

//字符串写文件 bufio
// 通过bufio写文件效率更高 先把内容写到内存缓冲区中
func writeFileString()  {
	if file, err := os.OpenFile("la01.text",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666); err!= nil{
		fmt.Printf("open file fail ,err: %s", err)
	} else {
		defer file.Close()
		fmt.Println("open file success!")
		writer := bufio.NewWriter(file) // 接收文件句柄的参数
		writer.WriteString("hello gogogo\n") // 需要显示添加换行符 先把内容写到内存缓冲区中
		writer.Flush() // 需要flush 强行把缓存区内容到磁盘文件
	}
}
// 关于文件的操作 删除 移动 重命名等

func fileOps()  {
	if err := os.MkdirAll("p1/p2/p3",0777); err != nil{ // 创建级联目录
		fmt.Printf("mkdir err:%s", err)
	}
	if _,err := os.Create("p1/p2/aabb.txt"); err != nil{
		fmt.Printf("touch file err :%s", err)
	}
	os.Create("p1/p2/remove.txt")//   创建文件
	os.Remove("p1/p2/remove.txt") //  删除文件
	os.Rename("p1/p2/aabb.txt","p1/p2/new.txt") //  重命名文件

	// 文件属性
	file ,_ := os.Open("test.txt")
	fmt.Println(file.Name()) // 文件名称
	info ,_:= file.Stat()
	fmt.Println(info.Name())  // 文件名称
	fmt.Println(info.Mode()) // 文件的mode
	fmt.Println(info.IsDir())  // 是否为目录
	fmt.Println(info.Size())  // 文件大小
	fmt.Println(info.Sys())
	fmt.Println(info.ModTime()) // 修改时间

}

// 遍历目录下的文件
//filepath.Join 目录连接 使用这个函数  可以根据操作系统选择连接目录的分隔符
func walkDir(path string) error {
	if fileInfo ,err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _,file := range fileInfo {
			fmt.Println(file.Name())
			if file.IsDir() {
				if err := walkDir(filepath.Join(path,file.Name())); err != nil{
					return err
				}
			}

		}
	}
	return nil
}

// 通过文件写 -记录日志
func fileLogger()  {
	if file,err := os.OpenFile("file.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666);err != nil {
		fmt.Printf("日志文件打开失败%s",err)
	} else {
		defer file.Close()
		logWriter := log.New(file,"[testLog]",log.Ldate|log.Ltime|log.LUTC)
		logWriter.Println("hello")
		logWriter.Println("hello Go !")
	}
}

// 通过go调用系统命令 直接执行shell 命令
//对系统性能没有要求

func shellCall()  {
	if cmdPath,err := exec.LookPath("df") ; err != nil { // 找命令的路径
		fmt.Println("找不到df命令")
	} else {
		fmt.Println(cmdPath)
		cmd:= exec.Command("df","-h")
		outPut,_ := cmd.Output()
		fmt.Println(string(outPut))
	}

	cmd := exec.Command("ls","-lrth","test.txt")
	if output, err := cmd.Output();err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(output) // 字节数组
		fmt.Println(string(output)) // 字节数组转换为字符串
	}
}

// 文件压缩 与 解压
//go中有很多压缩库可以用

func compressFile()  {
	//file,_ := os.Open("compress01.log")
}

func main()  {
	//stdFmt()
	//readFileByte()
	//readFileString()

	//writeFileByte()
	//writeFileString()
	//fileOps()

	walkDir("p1")
	//fileLogger()

	//shellCall()
}
