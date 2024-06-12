package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main ()  {
	timeOut := time.Second * 3

	srcDb,err := sql.Open("mysql","executor:DYVY3chVs8ph@tcp(10.89.139.33:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local&timeout=" + timeOut.String())
	if err!= nil{
		fmt.Printf("连接数据库报错 %v",err)
	}
	//fmt.Printf("db:%v\n",srcDb)

	defer srcDb.Close()


	rows , err := srcDb.Query("select SUBSTRING_INDEX(host,':',1) as ip from information_schema.processlist group by ip;")
	if err != nil {
		fmt.Printf("查询数据库报错 %v",err)
	}
	//fmt.Printf("rows:%v",rows)

	destDb, err := sql.Open("mysql","executor:DYVY3chVs8ph@tcp(10.89.181.40:3306)/auditsql?charset=utf8mb4&parseTime=True&loc=Local&timeout=" + timeOut.String())
	if err != nil{
		fmt.Printf("连接数据库报错 %v",err)
	}
	defer destDb.Close()

	stmt, err := destDb.Prepare("insert ignore into ip_tbl(host) values (?)")
	if err != nil{
		panic(err.Error())
	}

	for rows.Next() {
		var ip string
		rows.Scan(&ip)
		fmt.Println(ip)
		if _, err := stmt.Exec(ip); err != nil{
			panic(err.Error())
		}
	}


}