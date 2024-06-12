package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
	"time"
)



var (
	dbUser string = "executor"
	dbPwd string = "DYVY3chVs8ph"
	dbPort int = 3306
	dbName string = "auditsql"
	dbHost string = "10.89.181.40"
	today string = time.Now().Format("2006-01-02")
	dcUrl string = "https://im-dichat.xiaojukeji.com/api/hooks/incoming/be6e6130-eeec-44bc-9a3a-271232fc463d" // 生产环境DC不能发送消息 ，需要替换
	failCount int = 0
	successCount int = 0
	failReason string

	topTableRows string = "【表行数Top10】\ntableSchema   tableName   tableRows\n"
	topTableSizes string = "【表容量Top10】\ntableSchema   tableName   tableSize\n"
)

var tableSizeMap =  make(map[string]string)



func sendDcMsg(sendMsg string){
	jsonStr :=  fmt.Sprintf(`{"text":"%s"}`, sendMsg)
	fmt.Println(jsonStr)

	requestBody := strings.NewReader(jsonStr)
	//requestBody := bytes.NewBuffer([]byte(jsonStr))

	// 创建POST请求
	req, err := http.NewRequest("POST",dcUrl,requestBody)
	if err != nil {
		fmt.Println("创建请求失败", err)
		return
	}

	// 设置请求头的Content-Type为 application/json
	req.Header.Set("Content-Type","application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败：", err)
		return
	}
	defer  resp.Body.Close()

	// 处理响应
	fmt.Println("响应状态吗：",resp.Status)
}


func main() {
	// 集群元数据的数据源
	metaDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=2s", dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Println(metaDsn)
	metaDb, err := sql.Open("mysql", metaDsn)
	if err != nil {
		msg := fmt.Sprintf("【表容量统计巡检 %s】\n打开元数据库%s:%d失败\n 具体报错:%s",today,dbHost,dbPort,err)
		fmt.Println(msg)
		sendDcMsg(msg)
		return // 退出程序，因为无法打开数据库连接
	}
	defer metaDb.Close()
	// 不要巡检 SQL平台所有的数据库 ，
	metaSql := "select cluster_id, ip, port  from core_dbtree where cluster_id!=18140 and type='mysql' and environment='生产' and role='mrw' group by cluster_id, ip, port;"

	metaRows, err := metaDb.Query(metaSql)
	if err != nil {
		msg := fmt.Sprintf("【表容量统计巡检 %s】\n 查询元数据失败\n具体报错：%s",today,err)
		sendDcMsg(msg)
		return // 退出程序，因为查询失败
	}
	defer metaRows.Close()

	for metaRows.Next() {
		var clusterId int
		var ip string
		var port int
		if err := metaRows.Scan(&clusterId, &ip, &port); err != nil {
			fmt.Printf("Scan() 错误：%s\n", err)
			failCount ++
			continue // 继续下一次循环
		}
		fmt.Printf("------------------------ cluster_id:%d, ip:%s, port:%d ------------------------\n", clusterId, ip, port)
		// 业务库dataSource
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=2s", dbUser, dbPwd, ip, port, "mysql")
		fmt.Println(dataSourceName)
		db , err := sql.Open("mysql",dataSourceName)
		if err != nil {
			msg := fmt.Sprintf("打开业务实例%s:%d失败 ：%s",ip,port,err)
			fmt.Println(msg)
			mapKey := fmt.Sprintf("%s-%d",ip,port)
			tableSizeMap[mapKey] = msg
			failCount ++
			//sendDcMsg(msg)
			continue
		}


		tableSizeSql := "select table_schema,table_name,ENGINE,CREATE_TIME,UPDATE_TIME,TABLE_ROWS,round(sum(data_length)/1024/1024/1024,2) as data_size," +
			"round(sum(index_length)/1024/1024/1024,2) as index_size,round(sum(data_free)/1024/1024/1024,2) as data_free,round(sum(data_length+index_length)/1024/1024/1024,2) as table_size ,round(sum(data_free)/sum(data_length+index_length),2) as data_free_percentage from information_schema.tables where table_schema not in ('test','mysql','sys','information_schema','performance_schema')  group by table_schema,table_name;"
		rows ,err := db.Query(tableSizeSql)
		if err != nil {
			fmt.Printf("查询数据失败 %s\n", err)
			msg := fmt.Sprintf("查询业务实例 %s:%d 失败 ：%s",ip,port,err)
			mapKey := fmt.Sprintf("%s-%d",ip,port)
			tableSizeMap[mapKey] = msg
			failCount ++
			continue
		}

		for rows.Next()  {
			var tableSchema string
			var tableName string
			var Engine sql.NullString
			var createTime sql.NullString
			var updateTime sql.NullString
			var tableRows sql.NullInt64
			var dataSize  sql.NullFloat64
			var indexSize sql.NullFloat64
			var dataFree sql.NullFloat64
			var tableSize sql.NullFloat64
			var dataFreePercentage sql.NullFloat64
			//var createTimeStr string
			//var updateTimeStr string
			if err = rows.Scan(&tableSchema,&tableName,&Engine,&createTime,&updateTime,&tableRows,&dataSize,&indexSize,&dataFree,&tableSize,&dataFreePercentage);err != nil {
				fmt.Printf("Scan() 错误：%s\n", err)
				continue // 继续下一次循环
			}
			//if createTime.Valid {
			//	createTimeStr = createTime.String
			//}
			//if updateTime.Valid {
			//	updateTimeStr = updateTime.String\
			//}
			//fmt.Println(tableSchema,tableName,createTimeStr,updateTimeStr,tableRows,tableSize)

			insertSql := "replace into inspection_table_size (cluster_id,env,ip,port,table_schema,table_name,engine,create_time,update_time,table_rows,data_size,index_size,data_free,table_size,data_free_percentage,date) " +
				"values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

			_ , err := metaDb.Exec(insertSql,clusterId,"prod",ip,port,tableSchema,tableName,Engine,createTime,updateTime,tableRows,dataSize,indexSize,dataFree,tableSize,dataFreePercentage,today)

			if err != nil {
				fmt.Println("插入数据失败", err)
				return
			}
			//rowsAffected , err := res.RowsAffected()
			//if err != nil {
			//	fmt.Println("获取影响的行数失败",err)
			//	return
			//}
			//fmt.Printf("成功插入%d行数据\n",rowsAffected)


		}
		// 不使用defer ，因为是在for循环中
		db.Close()
		rows.Close()
		successCount ++
	}

	// 组装DC消息
	//fmt.Println("successCount",successCount)
	//fmt.Println("failCount",failCount)
	//
	//fmt.Println(tableSizeMap)
	// 失败原因
	for k,v := range tableSizeMap {
		failReason = fmt.Sprintf("%s\n%s:%s",failReason,k,v)
	}

	//fmt.Println(failReason)



	tableRowSql := fmt.Sprintf("select table_schema,table_name,table_rows  from  inspection_table_size where date='%s' order by table_rows  desc limit 10;",today)
	tableSizeSql := fmt.Sprintf(" select table_schema,table_name,data_size  from  inspection_table_size where date='%s' order by data_size desc limit 10;",today)

	tableRow,err := metaDb.Query(tableRowSql)
	if err != nil {
		fmt.Println("查询数据失败")
	}

	for tableRow.Next() {
		var tableSchema string
		var tableName string
		var tableRows int
		tableRow.Scan(&tableSchema,&tableName,&tableRows)
		topTableRows = fmt.Sprintf("%s %s %s %d\n",topTableRows,tableSchema,tableName,tableRows)
	}

	fmt.Println(topTableRows)

	tableSizes,err := metaDb.Query(tableSizeSql)
	if err != nil {
		fmt.Println("查询数据失败")
	}

	for tableSizes.Next() {
		var tableSchema string
		var tableName string
		var tableSize float64
		tableSizes.Scan(&tableSchema,&tableName,&tableSize)
		topTableSizes = fmt.Sprintf("%s %s %s %.1fG\n",topTableSizes,tableSchema,tableName,tableSize)
	}
	fmt.Println(topTableSizes)

	msg := fmt.Sprintf("【表容量巡检】\n巡检日期：%s\n巡检成功实例个数: %d\n巡检失败实例个数： %d\n失败原因： %s\n%s\n%s", today,successCount,failCount,failReason,topTableSizes,topTableRows)
	fmt.Println(msg)
	sendDcMsg(msg)

	if err := metaRows.Err(); err != nil {
		fmt.Printf("遍历结果时出错：%s\n", err)
	}
}