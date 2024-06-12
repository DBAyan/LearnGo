package main

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	timeFmt2 = "2006-01-02"
	timeFmt3 = "2006/01/02"

)

// 设置为全局变量 为什么不设置为 常量 因为 time.LoadLocation 返回两个参数
var loc *time.Location

func init()  {
	loc,_ = time.LoadLocation("Asia/Shanghai")
}

// time.Time 类型转换为字符串
func timeFmt() {
	begin := time.Now()
	fmt.Println(begin)
	fmt.Println(begin.Format(timeFormat))
	fmt.Println(begin.Format("2006-01-02"))
	fmt.Println(begin.Format("20060102"))
	fmt.Println(begin.Format("2006/01/02"))
}

func ticker()  {
	tc := time.NewTicker(3*time.Second)
	defer tc.Stop() //
	for i := 1;  i<6; i++ {
		<- tc.C // 管道阻塞3秒
		fmt.Println(time.Now().Unix())
	}
}

func timer()  {
	fmt.Println(time.Now().Unix())
	tm := time.NewTimer(time.Second * 3) // second 本质是duration类型
	defer tm.Stop()
	<- tm.C
	fmt.Println(time.Now().Unix())
	<- time.After(time.Second * 3)
	fmt.Println(time.Now().Unix())
}


func main()  {
	now := time.Now()
	fmt.Printf("type of now is %T,value is %v\n",now,now) // type of now is time.Time,value is 2022-11-27 18:18:19.114463 +0800 CST m=+0.000111334

	// 时间戳格式
	fmt.Println(now.Unix()) // 秒
	//fmt.Println(now.UnixMilli()) // 毫秒
	//fmt.Println(now.UnixMicro()) // 微秒
	fmt.Println(now.UnixNano()) // 纳秒


	// 时间差 since
	begin := time.Now()
	//time.Sleep(1 * time.Second)
	end := time.Now()
	useTIme2 := end.Sub(begin)
	fmt.Println(useTIme2)
	useTime := time.Since(begin)
	fmt.Println(useTime.Seconds())
	fmt.Println(useTime.Nanoseconds())

	// 八小时之后的时间 add
	dua := time.Duration(8 * time.Hour)
	future := begin.Add(dua)
	fmt.Println(future)

	// 获取时间的详细信息 比如星期几
	fmt.Println(begin.Weekday()) // 星期几 Sunday
	fmt.Println(begin.Year()) //  2022 年
	fmt.Println(begin.YearDay()) // 331 今年的第几天
	fmt.Println(begin.Month()) // November 月
	fmt.Println(begin.Month().String())
	fmt.Println(int(begin.Month())) // 月份的数字输出
	fmt.Println(begin.Day()) // 27 这个月的几号
	fmt.Println(begin.Date()) // 2022 November 27
	fmt.Println(begin.Hour())  //  时
	fmt.Println(begin.Minute()) //  分
	fmt.Println(begin.Second()) // 秒

	// 时间格式化 时间转字符串
	timeFmt()

	//时间格式化 字符串转时间
	// 不要使用这个方式
	if mybtd,err := time.Parse("2006-01-02","1992-05-23"); err == nil {
		fmt.Println(mybtd)
		fmt.Println(mybtd.Year())
		fmt.Println(mybtd.Month())
		fmt.Println(mybtd.Day())

	} else {
		fmt.Println(err)
	}

	// 时区相关 工作中使用这种方式 需要三个参数 (时间格式layout ，时间字符串，时区) ，返回是time.Time格式
	if mybtd,err := time.ParseInLocation(timeFmt2,"1992-05-23",loc); err == nil {
		fmt.Println(mybtd)
		fmt.Println(mybtd.Year())
		fmt.Println(mybtd.Month())
		fmt.Println(mybtd.Day())

	} else {
		fmt.Println(err)
	}

	// 定时执行任务 周期执行任务
	//模拟 cron
	ticker()

	// 3秒后执行一次
	timer()
}



