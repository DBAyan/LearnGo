package main

import (
	"fmt"
	"log"

	//"time"
)
import _ "github.com/go-sql-driver/mysql"
import  "github.com/jmoiron/sqlx"


func initDb(uName string, pwd string, ipAddr string, pt int, dName string) *sqlx.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", uName, pwd, ipAddr, pt, dName)
	fmt.Printf("dsn: %s\n", dsn)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect fail, datail is [%s]", err.Error())
	}
	return Db
}

func main()  {
	//var (
	//	username string = "dbadmin_01"
	//	password string = "9WJKXw8e"
	//	serverIP string = "10.179.251.236"
	//	port int = 3306
	//	dbname string = "mysql"
	//)

	//var Db *sqlx.DB = initDb(username,password,serverIP,port,dbname)
	//fmt.Println(Db)
	//defer Db.Close()
	//if Db == nil{
	//	fmt.Println("connect err")
	//}
	db, err := sqlx.Connect("mysql", "dbadmin_01:9WJKXw8e@tcp(172.20.2.224:3300)/mysql?charset=utf8mb4&parseTime=true&loc=Local")
	fmt.Println(err)
	fmt.Println(db)
	if err != nil {
		log.Println("数据库连接失败")
	}
	fmt.Println(db)
	type userInfo struct {
		userName string
		hostIp string
	}
	ui := userInfo{}
	rows ,err := db.Query("select user,host from mysql.user")
	fmt.Println(ui)
	if err != nil{
		fmt.Printf("查询失败", err)
	}
	fmt.Println(rows)
	for rows.Next(){
		err := rows.Scan(&ui.hostIp, &ui.hostIp)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", ui)
	}

}