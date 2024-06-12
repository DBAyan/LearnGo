package main

import (
	"fmt"
	"time"
)
import _ "github.com/go-sql-driver/mysql"
import  "github.com/jmoiron/sqlx"

var (
	username string = "dbadmin_01"
	password string = "9WJKXw8e"
	serverIP string = "10.179.251.236"
	port int = 3306
	dbname string = "mysql"
)

type selectInfo struct {
	user string
}



func connectMysql(uName string, pwd string, serverIp string,prt int,dName string) (Dd *sqlx.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",uName,pwd,serverIp,prt,dName)
	fmt.Printf("dsn:%s\n",dsn)
	Db,err := sqlx.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("mysql connect fail ,detail is [%s]\n", err.Error())
	}
	Db.SetConnMaxLifetime(time.Second * 100)
	Db.SetMaxIdleConns(5)
	Db.SetMaxOpenConns(10)
	fmt.Printf("db: %v",Dd)

	return Dd
}

func queryMultiRow(Db *sqlx.DB)  {
	sqlStr:= "select user,host  from mysql.user"
	var users selectInfo
	err := Db.Select(&users, sqlStr)
	if err != nil{
		fmt.Println("查询执行失败！err:",err)
		return
	}
}

//func main()  {
//	//dbConnect := connectMysql(username,password,serverIP,port,dbname)
//	//fmt.Println(dbConnect)
//	var Db *sqlx.DB = connectMysql(username,password,serverIP,port,dbname)
//	if Db == nil {
//		fmt.Println("init db fail!")
//	}
//	fmt.Println(Db)
//
//	queryMultiRow(Db)
//}


