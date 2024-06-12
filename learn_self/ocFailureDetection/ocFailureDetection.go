package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
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
)

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

	sendMsg := fmt.Sprintf("【故障发现阶段】\n故障类型：'%s'\n故障详细描述：'%s'\n实例类型：'%s'\n故障集群 ：'%s'\n故障集群别名：'%s'\n故障集群域名：'%s'\n故障实例：'%s:%s'\n新主库：'%s':%s\nDetected %s on %s. Affected replicas: %d",
		failureType,failureDescription,instanceType,failureCluster, failureClusterAlias,failureClusterDomain,failedHost,failedPort,successorHost,successorPort,failureType,failureCluster,countReplicas)
	fmt.Println(sendMsg)

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
