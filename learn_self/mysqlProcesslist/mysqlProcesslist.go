package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var (
	userName string = "dbadmin_01"
	password string = "9WJKXw8e"
	ipAddr string = "10.79.21.29"
	port int = 3306
	dbName string = "information_schema"
	id int
	user string
	host string

)

//type  processlist struct {
//	id int
//	user string
//	host string
//	db sql.NullString
//	command string
//	time int
//	state sql.NullString
//	info sql.NullString
//}



func main()  {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", userName, password, ipAddr, port, dbName)
	fmt.Println(dataSourceName)
	db, err := sql.Open("mysql",dataSourceName)
	if  err != nil {
		panic("打开数据库报错")
	}
	defer db.Close()


	rows , err := db.Query("show  processlist;")
	if err != nil {
		panic("查询报错")
	}

	defer  rows.Close()

	for rows.Next() {
		var (
			db sql.NullString
			command string
			time int
			state sql.NullString
			info sql.NullString
			dbStr string
			stateStr string
			infoStr string
		)

		if err := rows.Scan(&id, &user, &host, &db, &command, &time, &state, &info); err != nil {
			panic(err.Error())
		}
		if db.Valid {
			dbStr = db.String
		} else {
			dbStr = "NULL"
		}
		if state.Valid {
			stateStr = state.String
		} else {
			stateStr="NULL" }

		if info.Valid {
			infoStr = info.String
		} else {
			infoStr ="NULL"
		}

		fmt.Printf("ID: %d, User: %s, Host: %s, DB: %s, Command: %s, Time: %d, State: %s, Info: %s\n", id, user, host, dbStr, command, time, stateStr,infoStr)
		killSql := fmt.Sprintf("kill %d", id)
		fmt.Println(killSql)

	}


}
