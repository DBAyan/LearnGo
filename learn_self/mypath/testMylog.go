package main

import "fmt"
// 通一个模块下 ，改用该模块下包中的函数
import  "testlog1234/mylog" // 模块名/目录名 ，这样就可以使用该目录下的包中的函数 例如 mylog123.NewLogger

func main()  {
	myLogger,err := mylog123.NewLogger(mylog123.LogLevelDebug,true,true,true,true,"20230428.log")
	if err != nil{
		fmt.Printf("初始化日志报错：%v",err)
		return
	}
	defer myLogger.Close()
	myLogger.Debug("测试日志123")
	myLogger.Info("info 测试日志")
	myLogger.Warn("测试日志warn")
	myLogger.Error("测试日志error")
}
