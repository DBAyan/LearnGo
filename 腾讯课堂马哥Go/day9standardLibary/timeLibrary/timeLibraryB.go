package main

import (
	"fmt"
	"time"
)

// http://c.biancheng.net/view/5392.html

func main()  {
	var t1 time.Time
	fmt.Println(t1) // 0001-01-01 00:00:00 +0000 UTC

	// implicit assignment of unexported field 'wall' in time.Time literal T
	// ime 结构体中字段名都是小写 不能显示的赋值
	//t2 := time.Time{1,1,time.Local}
	//fmt.Println(t2)

	// 1.获取当前时间
	t2 := time.Now()
	fmt.Printf("type of t2 :%T\n", t2) // time.Time

	fmt.Println(t2) // 2022-12-13 18:28:24.3663 +0800 CST m=+0.000210835

	// 2. 获取指定时间
	t3 := time.Date(2022,12,13,19,30,24,0,time.Local) // 函数
	fmt.Println(t3) // 2022-12-13 19:30:24 +0800 CST

	//3 .Time  --> string 转换
	fmt.Println(t2.Format("2006-01-02 15:04:05")) // 2022-12-13 18:34:22
	fmt.Println(t2.Format("2006年01月02日 15时04分05秒")) // 2022年12月13日 18时36分04秒
	fmt.Println(t2.Format("2006/01/02 15:04:05")) // 2022/12/13 18:37:05

	// 4 .string --> time 转换
	/*
	time.Parse(模板,string) --> err,time.Time
	*/
	t3,err := time.Parse("2006-01-02 15:04:05","1992-05-23 12:30:31")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3) // 1992-05-23 12:30:31 +0000 UTC

	//5 . 根据当前时间 获取指定信息

	fmt.Printf("年：%d\n",t2.Year())
	fmt.Printf("月：%d\n",t2.Month())
	fmt.Printf("月：%s\n",t2.Month()) // 英文的月份
	fmt.Printf("日：%d\n",t2.Day())
	fmt.Printf("时：%d\n",t2.Hour())
	fmt.Printf("分：%d\n",t2.Minute())
	fmt.Printf("秒：%d\n",t2.Second())
	fmt.Printf("纳秒：%d\n",t2.Nanosecond())
	fmt.Printf("今年的第几天：%d\n",t2.YearDay())
	fmt.Printf("星期：%s\n",t2.Weekday())
	fmt.Printf("星期：%d\n",t2.Weekday()) // 英文的星期几
	fmt.Println(t2.Date())   // 2022 December 13  这里的Date()是 结构体的方法，需要作用在结构上



	// 6 时间戳
	// 距离 1970-01-01 00：00：00的时间差值
	fmt.Println(t2.Unix()) // 1670929824
	fmt.Println(t2.UnixNano()) // 1670929824 38804000

	// 7 时间间隔 加时间
	fmt.Println(t2)
	fmt.Println(t2.Add(time.Second * 60)) // 这里的 Second Hour 是常量
	fmt.Println(t2.Add(time.Hour * 24))

	// 8 Sleep
	fmt.Println("begin")
	time.Sleep(time.Second * 3) // 让程序睡眠3秒
	fmt.Println("over")

}
