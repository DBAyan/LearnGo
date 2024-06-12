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


	topDbSizes string = "【库容量Top10】\ndatabase dbSize\n"
	dbCountMsg string
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
	metaDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=5s", dbUser, dbPwd, dbHost, dbPort, dbName)
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
		msg := fmt.Sprintf("【库容量统计巡检】\n 查询元数据失败\n具体报错：%s",err)
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
			continue
		}

		dbSizeSql := "select table_schema,round(sum(data_length)/1024/1024/1024,2) as data_size,round(sum(index_length)/1024/1024/1024,2) as index_size,round(sum(data_free)/1024/1024/1024,2) as data_free,round(sum(data_length+index_length)/1024/1024/1024,2) as schema_size ,round(sum(data_free)/sum(data_length+index_length),2) as data_free_percentage, count(TABLE_NAME)  as table_count from information_schema.tables where table_schema not in ('test','mysql','sys','information_schema','performance_schema','orchestrator')  group by table_schema"
		//tableSizeSql := "select table_schema,table_name,ENGINE,CREATE_TIME,UPDATE_TIME,TABLE_ROWS,round(sum(data_length)/1024/1024/1024,2) as data_size," +
		//	"round(sum(index_length)/1024/1024/1024,2) as index_size,round(sum(data_free)/1024/1024/1024,2) as data_free,round(sum(data_length+index_length)/1024/1024/1024,2) as table_size ,round(sum(data_free)/sum(data_length+index_length),2) as data_free_percentage from information_schema.tables where table_schema not in ('test','mysql','sys','information_schema','performance_schema')  group by table_schema,table_name;"
		rows ,err := db.Query(dbSizeSql)
		if err != nil {
			fmt.Printf("查询数据失败 %s\n", err)
			msg := fmt.Sprintf("查询业务实例 %s:%d 失败 ：%s",ip,port,err)
			mapKey := fmt.Sprintf("%s-%d",ip,port)
			tableSizeMap[mapKey] = msg
			failCount ++
			continue
		}

		for rows.Next()  {
			var (
				tableSchema string
				dataSize sql.NullFloat64
				indexSize sql.NullFloat64
				dataFree sql.NullFloat64
				schemaSize sql.NullFloat64
				dataFreePercentage sql.NullFloat64
				tableCount sql.NullInt64
			)
			//var createTimeStr string
			//var updateTimeStr string
			if err = rows.Scan(&tableSchema,&dataSize,&indexSize,&dataFree,&schemaSize,&dataFreePercentage,&tableCount);err != nil {
				fmt.Printf("Scan() 错误：%s\n", err)
				continue // 继续下一次循环
			}

			insertSql := "replace into inspection_db_size (cluster_id,env,ip,port,table_schema,db_data_size,db_index_siz,db_data_free,db_table_size,db_data_free_percentage,table_count,date) " +
				"values (?,?,?,?,?,?,?,?,?,?,?,?)"

			_ , err := metaDb.Exec(insertSql,clusterId,"product",ip,port,tableSchema,dataSize,indexSize,dataFree,schemaSize,dataFreePercentage,tableCount,today)

			if err != nil {
				fmt.Println("插入数据失败", err)
				return
			}


		}
		// 不使用defer ，因为是在for循环中
		db.Close()
		rows.Close()
		successCount ++
	}

	// 失败原因
	for k,v := range tableSizeMap {
		failReason = fmt.Sprintf("%s\n%s:%s",failReason,k,v)
	}

	//fmt.Println(failReason)



	dbSizeSql := fmt.Sprintf("select table_schema,db_table_size from inspection_db_size where date='%s' order by db_table_size desc limit 10;",today)

	dbSizes,err := metaDb.Query(dbSizeSql)
	if err != nil {
		fmt.Println("查询数据失败")
		return
	}

	for dbSizes.Next() {
		var tableSchema string
		var dbSize float64
		dbSizes.Scan(&tableSchema,&dbSize)
		topDbSizes = fmt.Sprintf("%s %s %.1fG\n",topDbSizes,tableSchema,dbSize)
	}
	fmt.Println(topDbSizes)

	tableCountSql := fmt.Sprintf(" select count(*)  , sum(table_count) from inspection_db_size where date='%s';",today)
	tableCountRow,err := metaDb.Query(tableCountSql)
	if err != nil {
		fmt.Println("查询数据失败")
		return
	}

	for tableCountRow.Next() {
		var (
			dbCount int
			tableCount int
		)
		tableCountRow.Scan(&dbCount,&tableCount)
		dbCountMsg = fmt.Sprintf("【数量统计】\n 业务库总数:%d\n 业务表总数:%d\n",dbCount,tableCount)

	}


	msg := fmt.Sprintf("【库容量巡检】\n巡检日期：%s\n巡检成功实例个数: %d\n巡检失败实例个数：%d\n%s【失败原因】 %s\n%s\n", today,successCount,failCount,dbCountMsg,failReason,topDbSizes)
	fmt.Println(msg)
	sendDcMsg(msg)

	if err := metaRows.Err(); err != nil {
		fmt.Printf("遍历结果时出错：%s\n", err)
	}
}

