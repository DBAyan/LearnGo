package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//Api_key: EP
//Secure_key: di@ep$di
var (
	ApiKey string = "EP"
	secret string = "di@ep$di"
	sginUri string = "/lvs/outer/rs_list"
	rsListUri string="http://autoproxy.sys.xiaojukeji.com:8009/lvs/outer/rs_list"


)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))

}

func main()  {
	// 取当前时间戳 并转化为字符串
	currentTime := time.Now().Unix()
	currentTimeStamp := strconv.FormatInt(currentTime,10)
	fmt.Println(currentTimeStamp)
	fmt.Printf("%s\n",currentTimeStamp)
	//message = apikey_securekey_uri_timestamp
	//message := fmt.Sprintf("%s_%s_%s_%s", key, secret, uri, ts)

	message := fmt.Sprintf("%s_%s_%s_%s", ApiKey, secret, sginUri, currentTimeStamp)
	fmt.Println(message)

	rsListSign := ComputeHmac256(message, secret)
	// 加密签名
	fmt.Println(rsListSign)

	//	查询rs192.168.6.97:1097所属的vs。
	// /lvs/outer/rs_list?
	//ip=192.168.6.97&
	//port=1097&
	//key=sys_test&
	//sign=bukKafWC3fTtsbpzM0BvGftO3Kw2DW2QIhcw1UW2cEY&
	//timestamp=1575969344

	var rsIp string = "10.90.49.45"
	var port string="3307"
	url := fmt.Sprintf("%s?key=%s&sign=%s&timestamp=%s&ip=%s&port=%v",rsListUri,ApiKey,rsListSign,currentTimeStamp,rsIp,port)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp ,err := client.Do(req)
	if err !=nil {
		panic(err)
	}
	defer resp.Body.Close()
	body ,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%v", body)
	fmt.Println(string(body))

	responseBytes := []byte(body)
	type Response struct {
		Status string `json:"status"`
		Code   string `json:"code"`
		Reason []struct {
			IP  string `json:"ip"`
			PORT  string `json:"port"`
			Vip []struct {
				Cluster        string `json:"cluster"`
				Warehouse      string `json:"warehouse"`
				Product        string `json:"product"`
				IP             string `json:"ip"`
				Port           string `json:"port"`
				Protocol       string `json:"protocol"`
				Check          string `json:"check"`
				ConnectTimeout int    `json:"connect_timeout"`
				ConnectPort    string `json:"connect_port"`
				Weight         int    `json:"weight"`
				NbGetRetry     int    `json:"nb_get_retry"`
				LbAlgo         string `json:"lb_algo"`
			} `json:"vip"`
		} `json:"reason"`
	}

	// 解码 JSON 数据到 Response 结构体中
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		fmt.Println("解码 JSON 数据失败：", err)
		return
	}

	// 输出 Response 结构体中的 IP 字段值
	fmt.Println(response.Reason[0].IP)
	fmt.Println(response.Reason[0].PORT)
	fmt.Println(response.Status)
	fmt.Println(response.Code)
	fmt.Println(response.Reason[0].Vip[0].IP)
	fmt.Println(response.Reason[0].Vip[0].Port)
	//fmt.Printf("%T",response.Status)

	if response.Status == "suc" && response.Code == "000" {
		fmt.Println("请求成功")
	}

}