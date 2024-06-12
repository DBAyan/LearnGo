package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	remoteMasterAddr string
	version bool = false
	help bool = false
	port int = 3306
	innodbBufferPoolSize string
	serverId int
	baseDir string
	dataDir string
	readOnly int
	lowerCaseTableNames int
	gtidMode string
	onlyInit bool = false
	mysqlPackage string
	remoteFile string
)

func main()  {
	flag.BoolVar(&version,"version",false,"安装MySQL脚本版本")
	flag.BoolVar(&help,"help",false,"帮助信息")
	flag.IntVar(&port,"port",3306,"mysql端口")
	flag.StringVar(&innodbBufferPoolSize,"innodb_buffer_pool_size","1G","innodb_buffer_pool_size")
	flag.IntVar(&serverId,"server-id",0,"mysql server-id")
	flag.StringVar(&baseDir,"basedir","","mysql安装目录")
	flag.StringVar(&dataDir,"datadir","","mysql数据目录")
	flag.IntVar(&readOnly,"read_only",0,"是否只读,1-只读，0-读写")
	flag.IntVar(&lowerCaseTableNames,"lower_case_table_names",0,"表名列名大小写敏感，0-敏感，1-不敏感")
	flag.StringVar(&gtidMode,"gtid-mode","OFF","GTID模式，ON-开启，OFF-关闭")
	flag.BoolVar(&onlyInit,"only-init",false,"添加该参数后不会再安装mysql软件，只会进行初始化")

	flag.StringVar(&remoteMasterAddr,"remoteMasterAddr","","数据库服务器IP地址")
	flag.StringVar(&mysqlPackage,"mysqlPackage","","MySQL安装包路文件")
	flag.StringVar(&remoteFile,"remoteFile","","远程存放MySQL安装包的路径")
	flag.Parse()

	if version {
		fmt.Println("Remote Install MySQL Go Script v1.0")
		return
	}
	if help {
		flag.PrintDefaults()
		return

	}

	// 远程服务器地址、用户名和SSH私钥文件路径
	username := "root"
	privateKeyPath := "/root/.ssh/id_rsa"
	cmd := "echo 'SSH is Success!'"

	err, output := execRemoteCmd(username,privateKeyPath,remoteMasterAddr,cmd)

	if err != nil {
		fmt.Println("执行命令失败", err)
		return
	}
	fmt.Println("命令输出", string(output))

	remoteFilePath := filepath.Dir(remoteFile)
	fmt.Println(remoteFilePath)

	cmd = fmt.Sprintf("[ -d %s ] && echo 1 || echo 0", remoteFilePath)

	fmt.Println(cmd)

	err, output = execRemoteCmd(username,privateKeyPath,remoteMasterAddr,cmd)
	if err != nil {
		fmt.Println("执行检查目录的命令失败:", err)
	}

	fmt.Println("命令输出", string(output))
	fmt.Printf("%T",string(output))

	if strings.TrimRight(string(output),"\n") == "0" {
		fmt.Println("目录不存在")
		err, output := execRemoteCmd(username,privateKeyPath,remoteMasterAddr,fmt.Sprintf("mkdir -p %s",remoteFilePath))

		if err!= nil {
			fmt.Println("创建存放mysql安装包目录失败,",err)
			return
		}
		fmt.Println(string(output))
	}

	fmt.Println("目标端存放安装包目录已创建")



	err = copyFile(username,mysqlPackage,remoteFile,remoteMasterAddr)
	if err != nil {
		fmt.Println("文件传输失败：", err)
		return
	}


	err = checkPort(remoteMasterAddr,port)

	if err == nil {
		fmt.Println("端口被占用")
		return
	}

	if baseDir == "" || dataDir == "" {
		log.Fatalln("baseDir 和 dataDir不能设置为空")
	}

	// 创建目录
	checkAndCreatDir(username,privateKeyPath,remoteMasterAddr, baseDir)
	checkAndCreatDir(username,privateKeyPath,remoteMasterAddr, dataDir)




}

func execRemoteCmd(username,privateKeyPath,remoteAddr,cmd string) ( error, []byte) {
	// 读取私钥文件
	privateKey , err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("读取秘钥文件失败", err)
		return err,nil
	}

	// 解析私钥
	signer,err := ssh.ParsePrivateKey(privateKey)
	if  err != nil {
		fmt.Println("解析秘钥失败", err)
		return err,nil
	}

	// 创建SSH 客户端配置
	config := &ssh.ClientConfig{
		User:              username,
		Auth:              []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
		BannerCallback:    nil,
		ClientVersion:     "",
		HostKeyAlgorithms: nil,
		Timeout:           0,
	}

	// 建立SSH连接
	client ,err := ssh.Dial("tcp", remoteAddr+":22", config)
	if err != nil {
		fmt.Println("建立SSH连接失败", err)
		return err,nil
	}
	defer client.Close()

	// 创建ssh Session
	session,err := client.NewSession()
	if err!= nil{
		fmt.Println("创建SSH Session失败", err)

	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Println("执行命令失败")
		return err, nil
	}

	fmt.Println(string(output))
	return nil, output

}

func copyFile(username,localFilePath,remoteFilePath ,remoteAddr string) error {

	cmd := exec.Command("scp",localFilePath, username+"@"+remoteAddr+":"+remoteFilePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("拷贝文件失败：", err)
		return err
	}

	fmt.Println("拷贝文件成功")
	return nil
}


func checkPort(remoteAddr string, remotePort int)  error  {


	conn,err := net.DialTimeout("tcp",remoteAddr+":"+strconv.Itoa(remotePort),2*time.Second)

	if err == nil {
		fmt.Println("端口被占用")
		defer conn.Close()

	} else {
		fmt.Println("端口未被占用")
	}

	return err

}

func  checkAndCreatDir(username,privateKeyPath,remoteAddr, dir string) error  {
	cmd := fmt.Sprintf("[ -d %s ] && echo 1 || echo 0", dir)
	err , output := execRemoteCmd(username,privateKeyPath,remoteAddr,cmd)
	if err != nil {
		fmt.Println("执行检查目录的命令失败:", err)
	}

	if strings.TrimRight(string(output),"\n") == "0" {
		fmt.Println("目录不存在,创建目录")
		err, _ := execRemoteCmd(username,privateKeyPath,remoteMasterAddr,fmt.Sprintf("mkdir -p %s",dir))

		if err!= nil {
			fmt.Println("创建目录失败,",err)
			return err
		}
	}
	return nil

}

