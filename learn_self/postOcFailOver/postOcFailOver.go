package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//"echo 'failureType:{failureType},instanceType: {instanceType}, isMaster：{isMaster}, isCoMaster：{isCoMaster}, failureDescription：{failureDescription},
//command：{command}, failedHost：{failedHost}, failureCluster：{failureCluster}, failureClusterAlias：{failureClusterAlias}, failureClusterDomain：{failureClusterDomain},
//failedPort：{failedPort}, successorHost：{successorHost}, successorPort：{successorPort}, successorAlias：{successorAlias}, countReplicas：{countReplicas},
//replicaHosts：{replicaHosts},isDowntimed： {isDowntimed}, autoMasterRecovery：{autoMasterRecovery}, autoIntermediateMasterRecovery：{autoIntermediateMasterRecovery}' >> FailureDetection_var.log"


var (
	failureType string = ""
	instanceType string
	isMaster bool
	isCoMaster bool
	failureDescription string
	command string
	failedHost string
	failureCluster string
	failureClusterAlias string
	failureClusterDomain string
	failedPort string
	successorHost string
	successorPort string
	successorAlias string
	countReplicas int
	replicaHosts int
	isDowntimed bool
	autoMasterRecovery bool
	autoIntermediateMasterRecovery bool
	help bool=false
	dcUrl string = "https://api-kylin.intra.xiaojukeji.com/snitch_api_online_lb/hooks/1/incoming/8a52b19a-82c3-4dab-9c3f-6ecb7056e827" // 生产环境DC不能发送消息 ，需要替换

	// LVS 接口参数
	ApiKey  string = "EP"
	secret  string = "di@ep$di"

	sginUri string = "/lvs/outer/rs_update"
	vipListSgin string = "/lvs/outer/vip_list"
	//rsAddDelsginUri string = "/lvs/outer/rs_update"
	rsAddDelSgin string = "/lvs/outer/rs_add_del"
	vPoortCreteSgin string = "/lvs/outer/vport/create"

	// 不能在办公网访问LVS的接口了，但是在线上环境还可以
	vipListUrl      string = "http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/vip_list"
	rsUpdateUri     string = "http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/rs_update"
	rsAddDelUri     string = "http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/rs_add_del"
	vPortCreteUrl  string = "http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/vport/create"

)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))

}

func exitWithError(err error)  {

	fmt.Fprintf(os.Stderr,"error:%v",err)
	os.Exit(1)
}

func main(){
	flag.StringVar(&failureType,"failureType","","失败类型")
	flag.StringVar(&instanceType,"instanceType","","实例类型")
	flag.StringVar(&failureDescription,"failureDescription","","失败类型的具体描述")
	flag.StringVar(&command,"command","","命令")
	flag.StringVar(&failedHost,"failedHost","","故障实例HOST")
	flag.StringVar(&failedPort,"failedPort","","故障实例端口")
	flag.StringVar(&failureCluster,"failureCluster","","故障集群类型")
	flag.StringVar(&failureClusterAlias,"failureClusterAlias","","故障集群别名")
	flag.StringVar(&failureClusterDomain,"failureClusterDomain","","故障集群域名")
	flag.StringVar(&successorPort,"successorPort","","提升为主库的实例端口")
	flag.StringVar(&successorAlias,"successorAlias","","提升为主库的别名")
	flag.StringVar(&successorHost,"successorHost","","提升为主的实例HOST")
	flag.BoolVar(&isMaster,"isMaster",true,"是否为主库")
	flag.BoolVar(&isCoMaster,"isCoMaster",true,"是否为")
	flag.BoolVar(&isDowntimed,"isDowntimed",false,"是否设置为维护状态Downtimed")
	flag.BoolVar(&autoMasterRecovery,"autoMasterRecovery",true,"")
	flag.BoolVar(&autoIntermediateMasterRecovery,"autoIntermediateMasterRecovery",true," ")
	flag.IntVar(&countReplicas,"countReplicas",0,"")
	flag.IntVar(&replicaHosts,"replicaHosts",0,"")
	flag.BoolVar(&help,"help",false,"帮助信息")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
	fmt.Println(failureType)

	// 测试使用 实际是传入
	//failureClusterDomain := "10.88.151.158:5306"
	failureClusterDomainSlice := strings.Split(failureClusterDomain,":")
	vip := failureClusterDomainSlice[0]
	vPort := failureClusterDomainSlice[1]
	fmt.Println(vip)
	fmt.Println(vPort)


	//vipList(vip, vPort)
	err := rsAddDel(vip, vPort, "add")
	//err := vPortCreate(vip,vPort)
	if err != nil {
		sendMsg := fmt.Sprintf("【故障转移后阶段】\n故障类型：%s\n故障详细描述：'%s'\n实例类型：%s\n故障集群 ：%s\n故障集群别名：%s\n故障集群域名：%s\n故障实例：'%s:%s'\n新主库：'%s:%s'\n具体操作：新增vip %s:%s 新增后端RS %s:%s 失败\n失败原因：%s",
			failureType,failureDescription,instanceType,failureCluster, failureClusterAlias,failureClusterDomain,failedHost,failedPort,successorHost,successorPort,vip, vPort, successorHost, successorPort,err)
		sendDcMsg(sendMsg)
		// 以非 0 值退出
		exitWithError(fmt.Errorf("VIP %s:%s 新增后端RS %s:%s失败：%s",vip,vPort,failedHost,failedPort,err))
	}

	sendMsg := fmt.Sprintf("【故障转移后阶段】\n故障类型：%s\n故障详细描述：'%s'\n实例类型：%s\n故障集群 ：%s\n故障集群别名：%s\n故障集群域名：%s\n故障实例：'%s:%s'\n新主库：'%s:%s'\n具体操作：新增vip %s:%s 新增后端RS %s：%s",
		failureType,failureDescription,instanceType,failureCluster, failureClusterAlias,failureClusterDomain,failedHost,failedPort,successorHost,successorPort,vip, vPort, successorHost, successorPort)
	fmt.Println(sendMsg)
	// 发送DC消息
	sendDcMsg(sendMsg)

}

// 发送DC消息 接受字符串参数
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

func vipList(vip ,vPort string)  (cluster ,warehouse, product string,err error )  {

	// 获取当前时间
	currentTime := time.Now().Unix()
	currentTimeStamp := strconv.FormatInt(currentTime, 10)
	fmt.Println(currentTimeStamp)
	fmt.Printf("%s\n", currentTimeStamp)

	// 组装加密签名参数
	message := fmt.Sprintf("%s_%s_%s_%s", ApiKey, secret, vipListSgin, currentTimeStamp)
	fmt.Println(message)

	// 加密签名
	rsListSign := ComputeHmac256(message, secret)

	fmt.Println(rsListSign)


	url := fmt.Sprintf("%s?vip=%s&port=%d&key=%s&sign=%s&timestamp=%s",  vipListUrl, vip,vPort,ApiKey, rsListSign, currentTimeStamp)
	fmt.Println(url)

	req, err := http.NewRequest("GET",url, nil)
	if err != nil {
		fmt.Println("创建请求失败", err)
		return "","","" ,err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return "","","" ,err
	}

	defer  resp.Body.Close()

	// 处理响应
	fmt.Println("resp.Status（状态码，以及话术）：",resp.Status) // 状态码，以及话术 200 OK
	fmt.Println("resp.StatusCode(状态码):",resp.StatusCode) // 只有状态码 例如 200 ，400
	fmt.Println("resp.Proto(HTTP协议)：",resp.Proto) //  e.g. "HTTP/1.0"
	fmt.Println("resp.ContentLength", resp.ContentLength) // 内容长度
	fmt.Println("resp.Request（响应头）",resp.Header) // 响应头
	fmt.Println("resp.ProtoMajor:",resp.ProtoMajor)
	fmt.Println("resp.ProtoMinor:",resp.ProtoMinor)
	fmt.Println("resp.Body：",resp.Body) // 响应正文？
	fmt.Printf("resp.Body：类型 %T\n",resp.Body)
	fmt.Println("resp.Request",resp.Request)


	// 这样转换为[]byte类型
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("body 字节数组：%v\n",body)
	fmt.Printf("body类型：%T\n",body)

	if err != nil {
		fmt.Println(err)
		return "","","" ,err
	}
	// 转换为字符串
	fmt.Printf("Body 字符串%s\n",string(body))

	var data map[string]interface{}

	// 注意：这里传入的是data的地址 ，否则报错
	err = json.Unmarshal(body,&data)
	if err != nil {
		fmt.Println("解析body 为 json 失败：",err)
		return "","","" ,err
	}


	fmt.Println(data["status"])
	fmt.Printf("%T\n",data["status"])

	if data["status"] !="suc" {
		fmt.Println("vipList错误：:",string(body))
		err := errors.New(string(body))
		return "","","" ,err

	}

	status ,OK := data["status"].(string)
	if !OK {
		err:= fmt.Errorf("status类型断言失败")
		return "","","",err
	}
	fmt.Println("status:",status)

	reason ,OK := data["reason"].([]interface{})
	if !OK {
		err:= fmt.Errorf("reason 类型断言失败")
		return "","","",err
	}
	fmt.Println("reason：",reason)
	fmt.Printf("%T",reason)


	firstEle,OK := reason[0].(map[string]interface{})
	if !OK {
		err:= fmt.Errorf("reason firstEle 类型断言失败")
		return "","","",err
	}
	cluster ,OK = firstEle["cluster"].(string)
	if !OK {
		err:= fmt.Errorf(" firstEle  cluster 类型断言失败")
		return "","","",err
	}
	fmt.Println("cluster:", cluster)

	warehouse ,OK = firstEle["warehouse"].(string)
	if !OK {
		err:= fmt.Errorf(" firstEle  warehouse 类型断言失败")
		return "","","",err
	}
	fmt.Println("warehouse:",warehouse)

	product,OK  =  firstEle["product"].(string)
	if !OK {
		err:= fmt.Errorf(" firstEle  product 类型断言失败")
		return "","","",err
	}
	fmt.Println("product:", product)

	return cluster,warehouse,product,nil


}

func rsAddDel(vip ,vPort , optFlag string)  error {

	// 获取当前时间
	currentTime := time.Now().Unix()
	currentTimeStamp := strconv.FormatInt(currentTime, 10)
	fmt.Println(currentTimeStamp)
	fmt.Printf("%s\n", currentTimeStamp)

	// 组装加密签名参数
	message := fmt.Sprintf("%s_%s_%s_%s", ApiKey, secret, rsAddDelSgin, currentTimeStamp)
	fmt.Println(message)

	// 加密签名
	sign := ComputeHmac256(message, secret)

	fmt.Println(sign)

	url := fmt.Sprintf("%s?&key=%s&sign=%s&timestamp=%s",  rsAddDelUri, ApiKey, sign, currentTimeStamp)
	fmt.Println(url)

	cluster,warehouse,product ,err:= vipList(vip,vPort)
	if err != nil {
		fmt.Println(err)
		return  err
	}

	jsonStr := fmt.Sprintf(`{
    "cluster":"%s",
    "warehouse":"%s",
    "product":"%s",
    "productname": "001069",
    "vip":[
        {
            "ip":"%s",
            "port":"%s",
            "protocol":"TCP",
            "member":{
                "%s":[
                    {
                        "ip":"%s",
                        "port":"%s"
                    }
                ]
            }
        }
    ]
}
`,cluster,warehouse,product ,vip,vPort,optFlag, successorHost,successorPort)

	fmt.Println(jsonStr)

	payload := bytes.NewBuffer([]byte(jsonStr))


	//	payload := strings.NewReader(`{
	//    "cluster":"ilvs02_GZ-YS",
	//    "warehouse":"GZ-YS",
	//    "product":"001069",
	//    "productname": "001069",
	//    "vip":[
	//        {
	//            "ip":"10.88.151.158",
	//            "port":"5306",
	//            "protocol":"TCP",
	//            "member":{
	//                "del":[
	//                    {
	//                        "ip":"10.79.23.45",
	//                        "port":"5306"
	//                    }
	//                ]
	//            }
	//        }
	//    ]
	//}
	//`)

	fmt.Println(payload)
	fmt.Printf("%T\n", payload)


	req, err := http.NewRequest("POST",url, payload)
	if err != nil {
		fmt.Println("创建请求失败", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return err
	}

	defer  resp.Body.Close()

	// 处理响应
	fmt.Println("响应状态吗：",resp.Status)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		strBody := strings.ReplaceAll(string(body),"{","")
		strBody = strings.ReplaceAll(strBody,"}","")
		strBody = strings.ReplaceAll(strBody,"\"","")
		return errors.New(strBody)
	}

	return nil

}

func vPortCreate(vip ,vPort  string) error {

	// 获取当前时间
	currentTime := time.Now().Unix()
	currentTimeStamp := strconv.FormatInt(currentTime, 10)
	fmt.Println(currentTimeStamp)
	fmt.Printf("%s\n", currentTimeStamp)

	// 组装加密签名参数
	message := fmt.Sprintf("%s_%s_%s_%s", ApiKey, secret, vPoortCreteSgin, currentTimeStamp)
	fmt.Println(message)

	// 加密签名
	sign := ComputeHmac256(message, secret)

	fmt.Println(sign)

	url := fmt.Sprintf("%s?&key=%s&sign=%s&timestamp=%s",  vPortCreteUrl, ApiKey, sign, currentTimeStamp)
	fmt.Println(url)

	cluster,warehouse,product,err := vipList(vip,vPort)

	//"cluster": "ilvs_BJ-ZJY",
	//"warehouse": "BJ-ZJY",
	//"protocol": "TCP",
	//"product": "001069",

	jsonStr := fmt.Sprintf(`{
		"warehouse":"%s",
		"cluster":"%s",
		"product":"%s",
		"productname":"001069",
		"vport":[
				{
				"protocol":"TCP",
				"ip":"%s",
				"port":"%s",
				"persistence_timeout":900,
				"member":[
				{
				"ip":"%s",
				"port":"%s"
				}
				]
				}
			]
		}`,warehouse,cluster,product,vip,vPort,successorHost,successorPort)

	fmt.Println(jsonStr)
	payload := bytes.NewBuffer([]byte(jsonStr))
	req, err := http.NewRequest("POST",url, payload)
	if err != nil {
		fmt.Println("创建请求失败", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return err
	}

	defer  resp.Body.Close()

	// 处理响应
	fmt.Println("响应状态吗：",resp.Status)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	// 响应状态码不为 200  ，就可以确定请求失败？
	if resp.StatusCode != 200{
		// 处理body 中的大括号 和 双引号，将作为DC消息发送
		fmt.Println(string(body))
		strBody := strings.ReplaceAll(string(body),"{","")
		strBody = strings.ReplaceAll(strBody,"}","")
		strBody = strings.ReplaceAll(strBody,"\"","")
		return errors.New(strBody)
	}

	return nil

}


