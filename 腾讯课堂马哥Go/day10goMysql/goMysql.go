package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // sql: unknown driver "mysql" (forgotten import?)
	"time"
)

const timeLayout = "2006-01-02"
var (
	loc *time.Location
)

func init()  {
	loc,_ = time.LoadLocation("Asia/Shanghai")
}
// 插入

func insert(db *sql.DB)  {
	res,err := db.Exec("insert into t values(7,'yhh'),(8,'nico'),(9,'zs');")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	lastId,_ := res.LastInsertId()
	insertRwos ,_ := res.RowsAffected()
	//fmt.Println(lastId)
	//fmt.Println(insertRwos)
	fmt.Printf("插入%d行数据，当前id为%d",insertRwos,lastId)
}

func stmtInsert(db *sql.DB) {
	stmt,_ := db.Prepare("insert into t values(?,?,?),(?,?,?);")
	date1,_:= time.ParseInLocation(timeLayout,"2022-12-03",loc)
	date2,_:= time.ParseInLocation(timeLayout,"2018-05-15",loc)
	fmt.Printf("day of date1 is %d\n",date1.Local().Day())
	res ,err := stmt.Exec(11,"luffy", date1,22,"xiaojuan",date2)
	if err != nil {
		fmt.Println(err)
	}
	lastInsertId,_ := res.LastInsertId()
	fmt.Printf("after insert last id is %d\n", lastInsertId)
	rows,_ := res.RowsAffected()
	fmt.Printf("插入%d\n", rows)


}

// 更新操作
func update(db *sql.DB)  {
	res,_ := db.Exec("update t set name='yanhaihang' where name='yhh';")
	lastId ,_:= res.LastInsertId() // 更新为0
	updateRwos ,_ := res.RowsAffected()
	fmt.Printf("更新%d行数据，当前id为%d",updateRwos,lastId)

}

// 插入(覆盖)数据
func replace(db *sql.DB){

}

// 删除
func delete(db *sql.DB)  {
	res,_ := db.Exec("delete  from t where name='nico';")
	lastId ,_:= res.LastInsertId() // 为0，仅插入会返回最新的ID
	deleteRwos ,_ := res.RowsAffected()
	fmt.Printf("删除%d行数据，当前id为%d",deleteRwos,lastId)
}

// 查询
func query(db *sql.DB)  {
	rows,_ := db.Query("select id,name from t where name='yanhaihang';")
	for rows.Next() { //没有数据或发生error时返回false
		var id int
		var name string
		rows.Scan(&id,&name) //通过scan把db里的数据赋给go变量
		fmt.Printf("id :%d,name %s\n",id,name)
	}
}

func main()  {
	// 连接数据库 DNS格式 data source name  user:pwd@tcp(ip:port)/dbname?name=value&name=value
	// 简写 如果是本地mysql 端口为3306 则可以简写为 user:pwd/dbname
	// charset=utf8 charset=utf8mb4
	// parseTime=True 协助完成时间格式的解析与转换
	// 	db, err := sqlx.Connect("mysql", "dbadmin_01:9WJKXw8e@tcp(172.20.2.224:3300)/mysql?charset=utf8mb4&parseTime=true&loc=Local")
	db,err := sql.Open("mysql","dbadmin_01:9WJKXw8e@tcp(172.20.2.224:3300)/test?charset=utf8mb4")
	defer db.Close()
	//database.CheckError(err)
	if err != nil {
		//fmt.Println(err)
		fmt.Printf("数据库连接失败 %v", err)
	} else {
		//fmt.Println(db)
		fmt.Println("数据库连接成功！")
	}

	//insert(db)
	//update(db)
	//query(db)
	stmtInsert(db)
}
