package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

import _ "github.com/go-sql-driver/mysql"



type DaemonInfo struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	IP            string `json:"ip"`
	MGroup        string `json:"mgroup"`
	State         string `json:"state"`
	ApplyID       int    `json:"applyId"`
	UplineID      int    `json:"uplineId"`
	Msg           string `json:"msg"`
	LastModifyUser string `json:"lastModifyUser"`
	LastTime      string `json:"lastTime"`
	DDCloudRegion string `json:"ddcloud_region"`
}


var (
	username string = "executor"
	password string = "DYVY3chVs8ph"
	serverIP string = "10.89.181.40"
	port int = 3306
	dbname string = "auditsql"

	// 通过odin 接口获取得到的 数梦容器的IP
	odinIpList []string
	// 通过数据库获取到 数梦容器的IP
	dbIpList []string

	dcUrl  string = "https://im-dichat.xiaojukeji.com/api/hooks/incoming/be6e6130-eeec-44bc-9a3a-271232fc463d" // 生产环境DC不能发送消息 ，需要替换
)

func findSliceDiff(odinSlice,dbSlice [] string) []string  {
	var diff []string
	for _,odinEle := range odinSlice {
		found := false
		for _,dbEle := range dbSlice {
			if odinEle == dbEle {
				found = true
				break
			}
		}

		if !found {
			diff = append(diff,odinEle)
		}
	}

	return diff
}


func main()  {
	url := "http://tree.odin.intra.xiaojukeji.com/api/v1/ns/machine/list?ns=hnb-v.daemon.data-sync.datadream.didi.com"

	currentTime := time.Now()

	// 获取当前日期
	currentDate := currentTime.Format("2006-01-02")

	// 打印当前日期
	fmt.Println("Current Date:", currentDate)

	respone, err := http.Get(url)
	if err != nil {
		fmt.Println("发生错误", err)
		return
	}
	defer  respone.Body.Close()

	body, err := ioutil.ReadAll(respone.Body)

	// *http.gzipReader
	fmt.Printf("%T\n", respone.Body)

	fmt.Println("响应体：",string(body))

	var daemonInfoList []DaemonInfo

	err = json.Unmarshal(body,&daemonInfoList)
	if err != nil {
		fmt.Println("解析json发生错误：", err)
		return
	}
	for _, daemonInfo := range daemonInfoList {
		fmt.Printf("ID: %d, Name: %s, IP: %s, State: %s\n", daemonInfo.ID, daemonInfo.Name, daemonInfo.IP, daemonInfo.State)
		odinIpList = append(odinIpList, daemonInfo.IP)
	}

	fmt.Println(odinIpList)
	fmt.Println(odinIpList[0])

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",username,password,serverIP,port,dbname)
	db, err := sql.Open("mysql",dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select  ip  from  odin_ip")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ip string
		err := rows.Scan(&ip)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ip)
		dbIpList = append(dbIpList,ip)
	}

	fmt.Println(dbIpList)



	odinMore := findSliceDiff(odinIpList, dbIpList)
	fmt.Println(odinMore)
	fmt.Println("数梦odin多出IP ",odinMore)

	dbMore := findSliceDiff(dbIpList,odinIpList)
	fmt.Println(dbMore)

	var sendMsg string
	if len(odinMore) >0  || len(dbMore) >0  {
		sendMsg = fmt.Sprintf("【odin数梦IP与白名单比对巡检】\n巡检日期：%s\n数梦odin多出IP :%s,\ndb多出IP:%s ",currentDate,odinMore,dbMore)
		fmt.Println(sendMsg)
	} else {
		sendMsg = fmt.Sprintf("【odin数梦IP与白名单比对巡检】\n巡检日期：%s\n数梦odin IP 与 数据库白名单对比无差异",currentDate)
		fmt.Println(sendMsg)

	}






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

	deleteSQL := "delete from odin_ip"
	_, err  = db.Exec(deleteSQL)
	if err != nil{
		log.Fatal(err)
	}

	insertSQL := "insert into odin_ip (ip,host) values (?,?)"
	for _,daemonInfo := range daemonInfoList {
		_,err = db.Exec(insertSQL,daemonInfo.IP,daemonInfo.Name)
		if err!= nil {
			log.Fatal(err)
		}
	}
}