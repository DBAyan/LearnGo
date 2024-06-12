package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	slaveHost string
	masterHost string
	slavePort int
	user string
	password string
	database string
	masterServerId int
	help bool = false
)

func main()  {
	flag.StringVar(&slaveHost,"slaveHost","","从库的host")
	flag.StringVar(&masterHost,"masterHost","","主库的host")
	flag.IntVar(&slavePort,"slavePort",3306,"从库端口")
	flag.StringVar(&user,"user","","用户")
	flag.StringVar(&password,"password","","密码")
	flag.StringVar(&database,"database","","存放heartbeat表的数据库")
	flag.IntVar(&masterServerId,"masterServerId",0,"主库的server-id")
	flag.BoolVar(&help,"help",false,"帮助信息")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	slaveLagTicker := time.Tick(3* time.Second)

	for {
		<- slaveLagTicker
		go func() {
			// 组装操作系统命令
			heartBeatCmd := fmt.Sprintf("/usr/bin/pt-heartbeat --host='%s' --port=%d --user='%s' --password='%s' --database='%s' --master-server-id=%d --check",
				slaveHost,slavePort,user,password,database,masterServerId)



			fmt.Println(heartBeatCmd)

			// 执行操作系统命令并接受返回值

			cmd := exec.Command("/usr/bin/pt-heartbeat",
				fmt.Sprintf("--host=%s",slaveHost),
				fmt.Sprintf("--port=%d",slavePort),
				fmt.Sprintf("--user=%s",user),
				fmt.Sprintf("--password=%s",password),
				fmt.Sprintf("--database=%s",database),
				fmt.Sprintf("--master-server-id=%d",masterServerId),"--check")

			output ,err := cmd.CombinedOutput()
			if err != nil{
				fmt.Printf("执行命令%s报错 %s",heartBeatCmd, err)
				fmt.Println(string(output))
				return
			}
			fmt.Println(output)
			slaveLag := strings.TrimRight(string(output),"\n")
			fmt.Println(slaveLag)
			fmt.Printf("%T",slaveLag)

			slaveLagfloat, err := strconv.ParseFloat(slaveLag, 64)

			if err != nil {
				fmt.Printf("字符串转换小数失败%s",err)
			}

			fmt.Printf("%f",slaveLagfloat)

			// 执行操作系统命令并接受的第二种方式
			//cmd = exec.Command(heartBeatCmd)
			//cmd.Stderr = os.Stderr
			//cmd.Stdout = os.Stdout
			//err = cmd.Run()
			//if err != nil {
			//	fmt.Printf("执行命令%s报错 %s",heartBeatCmd, err)
			//	return
			//}
			//fmt.Println(cmd.Stderr)
			//fmt.Println(cmd.Stdout)
			//exitCode := cmd.ProcessState.ExitCode()
			//fmt.Println(exitCode)


			// 写入数据库
			// "username:password@tcp(127.0.0.1:3306)/dbname"
			masterUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",user,password,masterHost,slavePort,database)
			fmt.Println(masterUrl)
			db,err := sql.Open("mysql",masterUrl)

			if err != nil{
				log.Fatalf("连接数据库报错，%s",err)
			}

			defer db.Close()

			insertSql := fmt.Sprintf("replace  into slave_lag (host,lag) values ('%s',%f)",slaveHost,slaveLagfloat)
			fmt.Println(insertSql)
			//insertSql := "replace  into slave_lag (host,lag) values (?,?)"
			_, err = db.Exec(insertSql)

			if err != nil {
				log.Fatalf("插入数据库失败%s",err)
			}

		}()
	}

}
