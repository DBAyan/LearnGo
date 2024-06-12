package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	command                string
	origMasterIsNewSlave   bool
	origMasterHost         string
	origMasterIP           string
	origMasterPort         int
	origMasterUser         string
	origMasterPassword     string
	origMasterSSHUser      string
	newMasterHost          string
	newMasterIP            string
	newMasterPort          int
	newMasterUser          string
	newMasterPassword      string
	newMasterSSHUser       string
	runningIntervalSeconds float64 = 0.1
)

const (
	timeFormat = "2006-01-02 15:04:05.000000"
)

func init() {
	flag.StringVar(&command, "command", "", "The command to run")
	flag.BoolVar(&origMasterIsNewSlave, "orig_master_is_new_slave", false, "Whether the original master is now a new slave")
	flag.StringVar(&origMasterHost, "orig_master_host", "", "The hostname of the original master")
	flag.StringVar(&origMasterIP, "orig_master_ip", "", "The IP address of the original master")
	flag.IntVar(&origMasterPort, "orig_master_port", 3306, "The port number of the original master")
	flag.StringVar(&origMasterUser, "orig_master_user", "", "The username to connect to the original master")
	flag.StringVar(&origMasterPassword, "orig_master_password", "", "The password to connect to the original master")
	flag.StringVar(&origMasterSSHUser, "orig_master_ssh_user", "", "The SSH username to connect to the original master")
	flag.StringVar(&newMasterHost, "new_master_host", "", "The hostname of the new master")
	flag.StringVar(&newMasterIP, "new_master_ip", "", "The IP address of the new master")
	flag.IntVar(&newMasterPort, "new_master_port", 3306, "The port number of the new master")
	flag.StringVar(&newMasterUser, "new_master_user", "", "The username to connect to the new master")
	flag.StringVar(&newMasterPassword, "new_master_password", "", "The password to connect to the new master")
	flag.StringVar(&newMasterSSHUser, "new_master_ssh_user", "", "The SSH username to connect to the new master")
}

func main() {
	flag.Parse()

	switch command {
	case "stop":
		// Set read_only=1 on the new master
		newMasterDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/", newMasterUser, newMasterPassword, newMasterIP, newMasterPort)
		newMasterDB, err := sql.Open("mysql", newMasterDSN)
		if err != nil {
			log.Fatalf("Failed to connect to new master: %v", err)
		}
		defer newMasterDB.Close()
		// 1. Set read_only= 1 on the new master
		_, err = newMasterDB.Exec("SET GLOBAL read_only=1")
		if err != nil {
			log.Fatalf("Failed to set read_only on new master: %v", err)
		}
		log.Println("Set read_only on the new master")

		// Connect to the orig master and drop application user
		origMasterDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/", origMasterUser, origMasterPassword, origMasterIP, origMasterPort)

		origMasterDB, err := sql.Open("mysql", origMasterDSN)
		if err != nil {
			log.Fatalf("Failed to connect to orig master(%v:%v): %v", origMasterIP, origMasterPort, err)
		}
		defer origMasterDB.Close()

		// 3. Set read_only= 1 on the current master
		_, err = origMasterDB.Exec("SET GLOBAL read_only=1")
		if err != nil {
			log.Fatalf("Failed to set read_only on orig master(%v:%v): %v", origMasterIP, origMasterPort, err)
		}
		log.Printf("Set read_only on the orig master(%v:%v)", origMasterIP, origMasterPort)

		var connID int64
		err = origMasterDB.QueryRow("SELECT CONNECTION_ID()").Scan(&connID)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		log.Printf("Connection ID: %d\n", connID)
		// TODO 记得替换该VIP 现在是测试VIP
		// 4  down orig master interface VIP
		stopVipCmd := "ifconfig  eth0:0 10.79.23.3 netmask 255.255.255.128 down"

		err = stopVIP(origMasterIP,origMasterSSHUser,stopVipCmd)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Printf("stop VIP on orig master(%v:%v) success!", origMasterIP, origMasterPort)

		// 5 kill session

		////threads := []map[string]interface{}{}
		//
		//threads := getThreadsUtil(origMasterDB,connID,0,0)
		//log.Println(threads)
		//for _,threadItem := range  threads{
		//	log.Println(threadItem)
		os.Exit(0)
	case "start":
		rsAddDel()
		os.Exit(0)
	//case "status":
	//	log.Println("test status")

	}
}

func getThreadsUtil(db *sql.DB, myConnectionID int64, runningTimeThreshold int, t int) []map[string]interface{} {
	if runningTimeThreshold == 0 {
		runningTimeThreshold = -1
	}
	if t == 0 {
		t = -1
	}

	threads := []map[string]interface{}{}

	rows, err := db.Query("SHOW PROCESSLIST")
	if err != nil {
		log.Fatal(err)
	}


	defer rows.Close()

	for rows.Next() {
		var id int64
		var user string
		var host string
		var db string
		var command string
		var state string
		var queryTime int64
		var info string

		err = rows.Scan(&id, &user, &host, &db, &command, &state, &queryTime, &info)
		if err != nil {
			log.Fatal(err)
		}

		if myConnectionID == id {
			continue
		}
		if queryTime >= int64(runningTimeThreshold) {
			continue
		}
		if command == "Binlog Dump" {
			continue
		}
		if user == "system user" {
			continue
		}
		if command == "Sleep" && queryTime >= 1 {
			continue
		}

		if t >= 1 {
			if command == "Sleep" || command == "Connect" {
				continue
			}
		}

		if t >= 2 {
			if strings.HasPrefix(strings.ToLower(info), "select") || strings.HasPrefix(strings.ToLower(info), "show") {
				continue
			}
		}

		thread := map[string]interface{}{
			"Id":      id,
			"User":    user,
			"Host":    host,
			"Db":      db,
			"Command": command,
			"State":   state,
			"Time":    queryTime,
			"Info":    info,
		}

		threads = append(threads, thread)
	}

	return threads
}


func startVIP(newMasterHost string, sshUser string, sshStartVIP string) error {
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", sshUser, newMasterHost), sshStartVIP)
	return cmd.Run()
}

func stopVIP(origMasterHost string, sshUser string, sshStopVIP string) error {
	if sshUser == "" {
		return nil
	}
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", sshUser, origMasterHost), sshStopVIP)
	return cmd.Run()
}

//Api_key: EP
//Secure_key: di@ep$di
var (
	ApiKey  string = "EP"
	secret  string = "di@ep$di"
	sginUri string = "/lvs/outer/rs_update"

	rsUpdateUri     string = "http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/rs_update"
	rsAddDelUri     string = "http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/rs_add_del"
	rsAddDelsginUri string = "/lvs/outer/rs_update"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))

}


func updateRsConfig() {
	// 取当前时间戳 并转化为字符串
	currentTime := time.Now().Unix()
	currentTimeStamp := strconv.FormatInt(currentTime, 10)
	fmt.Println(currentTimeStamp)
	fmt.Printf("%s\n", currentTimeStamp)

	message := fmt.Sprintf("%s_%s_%s_%s", ApiKey, secret, sginUri, currentTimeStamp)
	fmt.Println(message)

	rsListSign := ComputeHmac256(message, secret)
	// 加密签名
	fmt.Println(rsListSign)

	url := fmt.Sprintf("%s?key=%s&sign=%s&timestamp=%s", rsUpdateUri, ApiKey, rsListSign, currentTimeStamp)
	fmt.Println(url)

	//type MemberDetail struct {
	//	Cluster            string `json:"cluster,omitempty"`
	//	Warehouse          string `json:"warehouse,omitempty"`
	//	Product            string `json:"product,omitempty"`
	//	ProductName        string `json:"productname,omitempty"`
	//	Ip                 string `json:"ip,omitempty"`
	//	Port               string `json:"port,omitempty"`
	//	Protocol           string `json:"protocol,omitempty"`
	//	Checkrs            string `json:"check,omitempty"`
	//	Path               string `json:"path,omitempty"`
	//	Connect_timeout    int    `json:"connect_timeout,omitempty"`
	//	Connect_port       string `json:"connect_port,omitempty"`
	//	Weight             int    `json:"weight,omitempty"`
	//	Nb_get_retry       int    `json:"nb_get_retry,omitempty"`       /* dgw连接rs的重试次数 */
	//	Inhibit_on_failure bool   `json:"inhibit_on_failure,omitempty"` /* 以前给LVS的参数,已废除 */
	//	Algorithm          string `gorm:"column:lb_algo" json:"lb_algo,omitempty"`
	//	PortRange          string `gorm:"column:port_range" json:"portRange,omitempty"` /* port范围 */
	//
	//	PortInt     *uint32 `json:"-"`
	//	ConnectPort *uint32 `json:"-"`
	//}
	//
	//type VipReqAPI struct {
	//	Ip        string `json:"ip"`
	//	Port      string `json:"port"`
	//	Protocol  string `json:"protocol"`
	//	PortRange string `gorm:"column:port_range" json:"portRange,omitempty"` /* port范围 */
	//}
	//
	//type MemUP struct { /* 一个rs可以挂载到多个vs下,故目测表示为member中的每个rs都属于Vip中指定的vs下 */
	//	Vip    []VipReqAPI
	//	Member []MemberDetail
	//}
	//
	//type MemberUpdateReqAPI struct {
	//	Cluster     string
	//	Warehouse   string
	//	Product     string
	//	ProductName string `json:"productname,omitempty"`
	//	Member      []MemUP
	//}

	// "reason":[{"ip":"10.90.49.45","port":"3307","vip":[{"cluster":"ilvs02_GZ-YS","warehouse":"GZ-YS","product":"001069","ip":"10.88.151.136","port":"3306",
	//"protocol":"TCP","check":"tcp","connect_timeout":5,"connect_port":"3307","weight":100,"nb_get_retry":3,"lb_algo":"rr"}]}]}

	payload := strings.NewReader(`{
    "warehouse": "GZ-YS",
    "cluster": "ilvs02_GZ-YS",
    "product": "001069",
	"productname": "001069",
    "member" : [
        {
            "vip": [
                {
                    "ip": "10.88.151.136",
                    "port": "3306",
                    "protocol": "tcp"
                }
            ],
            "member": [
                {
                    "ip": "10.90.49.44",
                    "port": "3306",
                    "weight": 100
                }
            ]
        }
    ]
}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", rsUpdateUri, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}

func rsAddDel() {
	// 取当前时间戳 并转化为字符串
	currentTime := time.Now().Unix()
	currentTimeStamp := strconv.FormatInt(currentTime, 10)
	//fmt.Println(currentTimeStamp)
	//fmt.Printf("%s\n", currentTimeStamp)

	message := fmt.Sprintf("%s_%s_%s_%s", ApiKey, secret, rsAddDelsginUri, currentTimeStamp)
	//fmt.Println(message)

	rsAddDelSign := ComputeHmac256(message, secret)
	// 加密签名
	//fmt.Println(rsAddDelSign)

	url := fmt.Sprintf("%s?key=%s&sign=%s&timestamp=%s", rsAddDelUri, ApiKey, rsAddDelSign, currentTimeStamp)
	log.Println(url)
	// "reason":[{"ip":"10.90.49.45","port":"3307","vip":[{"cluster":"ilvs02_GZ-YS","warehouse":"GZ-YS","product":"001069","ip":"10.88.151.136","port":"3306",
	//"protocol":"TCP","check":"tcp","connect_timeout":5,"connect_port":"3307","weight":100,"nb_get_retry":3,"lb_algo":"rr"}]}]}
	origMasterPortInt := strconv.Itoa(origMasterPort)
	newMasterPortInt := strconv.Itoa(newMasterPort)
	jsonStr := fmt.Sprintf("{\n    \"cluster\":\"ilvs02_GZ-YS\",\n    \"warehouse\":\"GZ-YS\",\n    \"product\":\"001069\",\n    \"productname\": \"001069\",\n    \"vip\":[\n        {\n            \"ip\":\"10.88.151.136\",\n            \"port\":\"3306\",\n            \"protocol\":\"TCP\",\n            \"member\":{\n                \"add\":[\n                    {\n                        \"ip\":\"%s\",\n                        \"port\":\"%s\"\n                    }\n                ],\n                \"del\":[\n                    {\n                        \"ip\":\"%s\",\n                        \"port\":\"%s\"\n                    }\n                ]\n            }\n        }\n    ]\n}",newMasterIP,newMasterPortInt,origMasterIP,origMasterPortInt)
	log.Println("调用LVS接口修改VIP后端real server")
	log.Println(jsonStr)
	payload := strings.NewReader(jsonStr)

	client := &http.Client{}
	req, err := http.NewRequest("POST", rsAddDelUri, payload)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(body))
	type Response struct {
		Status string `json:"status"`
		VIP    string `json:"vip"`
		Reason string `json:"reason"`
	}

	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.Status)

	if response.Status=="suc" {
		log.Printf("VIP切换成功，VIP后端RS由 %v:%v --> %v:%v",origMasterIP,origMasterPort,newMasterIP,newMasterPort)
	} else {
		log.Println("VIP切换失败")
	}

}
