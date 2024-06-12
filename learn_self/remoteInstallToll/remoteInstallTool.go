package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	remoteMasterAddr string
	remoteSlaveAddr string
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

// ANSI颜色转义码
const (
	Black   = "\u001b[30m"
	Red     = "\u001b[31m"
	Green   = "\u001b[32m"
	Yellow  = "\u001b[33m"
	Blue    = "\u001b[34m"
	Magenta = "\u001b[35m"
	Cyan    = "\u001b[36m"
	White   = "\u001b[37m"
	Reset   = "\u001b[0m"
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
	flag.StringVar(&remoteSlaveAddr,"remoteSlaveAddr","","从实例IP地址，可以传入多个，用逗号分割。例如slaveIP1,slaveIP2")
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


	// 安装主库
	//fmt.Printf("%s开始部署主库%s%s",Green,remoteMasterAddr,Reset)
	//installMydumper(remoteMasterAddr)
	if remoteSlaveAddr != "" {
		slaveAddrs := strings.Split(remoteSlaveAddr,",")
		for _,slaveAddr :=  range slaveAddrs {
			fmt.Printf("%s开始部署服务器%s\n%s",Green,slaveAddr,Reset)
			installMydumper(slaveAddr)
		}

	} else {
		log.Println("从库IP为空，为单实例模式")
	}

}



func createSshClient(remoteAddr string) ( error, *ssh.Client) {

	// 远程服务器地址、用户名和SSH私钥文件路径
	username := "root"
	privateKeyPath := "/root/.ssh/id_rsa"


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
	}

	// 建立SSH连接
	client ,err := ssh.Dial("tcp", remoteAddr+":22", config)
	if err != nil {
		fmt.Println("建立SSH连接失败", err)
		return err,nil
	}

	return nil,client

}

func execCmd(client *ssh.Client,cmd string) (error, string) {
	// 创建ssh Session
	session,err := client.NewSession()
	if err!= nil{
		fmt.Println("创建SSH Session失败", err)

	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Printf("执行命令失败%s",err)
		fmt.Println(string(output))
		return err,string(output)
	}

	fmt.Println(string(output))
	return nil, string(output)
}



func copyFile(username,localFilePath,remoteFilePath ,remoteAddr string) error {

	cmd := exec.Command("scp",localFilePath, username+"@"+remoteAddr+":"+remoteFilePath)
	err := cmd.Run()
	if err != nil {
		log.Printf("传输文件%s失败：%s", localFilePath,err)
		return err
	}

	fmt.Printf("传输文件%s成功",localFilePath)
	return nil
}

func installMydumper(remoteAddr string)  {

	err, client := createSshClient(remoteAddr)
	if err != nil {
		fmt.Printf("创建client失败：%s",err)
		return
	}
	fmt.Printf("%s创建client成功",remoteAddr)

	defer client.Close()

	remoteFilePath := filepath.Dir(remoteFile)
	fmt.Println(remoteFilePath)

	checkAndCreatDir(client,remoteFilePath)
	if err := copyFile("root",mysqlPackage,remoteFile,remoteAddr);err != nil {
		fmt.Printf("传输安装包失败：%s",err)
	}
	cmd := fmt.Sprintf("yum install -y %s",remoteFile)
	err, output := execCmd(client,cmd)
	fmt.Println(output)

	if err != nil{
		fmt.Printf("执行安装mydumper包失败，%s",err)
	}

	fmt.Println("安装mydumper软件包成功，")

}

func checkPort(remoteAddr string, remotePort int)  error  {

	conn,err := net.DialTimeout("tcp",remoteAddr+":"+strconv.Itoa(remotePort),2*time.Second)

	if err == nil {
		log.Printf("端口%d被占用", port)
		defer conn.Close()

	} else {
		log.Printf("端口%d未被占用",port)
	}

	return err

}

func checkDirExists(client *ssh.Client, dir string) bool {
	cmd := fmt.Sprintf("[ -d %s ] && echo 1 || echo 0", dir)

	err, output := execCmd(client,cmd)

	if err != nil {
		fmt.Println("执行检查目录的命令失败:", err)
	}

	if strings.TrimRight(output,"\n") == "0" {
		fmt.Printf("目录%s不存在",dir)
		return false
	}
	return true

}

func  checkAndCreatDir(client *ssh.Client, dir string) error  {
	cmd := fmt.Sprintf("[ -d %s ] && echo 1 || echo 0", dir)

	err, output := execCmd(client,cmd)

	if err != nil {
		fmt.Println("执行检查目录的命令失败:", err)
		return err
	}

	if strings.TrimRight(output,"\n") == "0" {
		fmt.Printf("目录%s不存在,创建目录",dir)
		err, _ := execCmd(client,fmt.Sprintf("mkdir -p %s",dir))

		if err!= nil {
			fmt.Printf("创建目录%s失败:%s",dir,err)
			return err
		}
	} else {
		log.Printf("目录%s已存在，终止执行",dir)
		//err := fmt.Errorf("目录%s已存在,终止执行",dir)
		return nil
	}
	return nil

}


func genServerId(remoteAddr string, port int) (int, error) {
	parts := strings.Split(remoteAddr,".")
	lastTwoParts :=  parts[len(parts)-2:]
	serverId := strings.Join(lastTwoParts,"") + strconv.Itoa(port)
	fmt.Println(serverId)

	return strconv.Atoi(serverId)
}


// 需要接受一个参数  ，否则使用了传入的参数
func createConfigFile(dataDirFull string) (error , string){
	// 配置文件模板
	mysqlConfig := fmt.Sprintf(
		"[client]\nport            = %v\nsocket            = %v/run/mysql.sock\n\n# The MySQL server\n[mysqld]\n#########Basic##################\nexplicit_defaults_for_timestamp=true\nport            = %v\nuser            = mysql\nbasedir         = %v\ndatadir         = %v/data\ntmpdir          = %v/tmp\npid-file        = %v/run/mysql.pid\nsocket            = %v/run/mysql.sock\n#skip-grant-tables\n#character set\ncharacter_set_server = utf8mb4\nopen_files_limit = 65535\nback_log = 500\n#event_scheduler = ON\nlower_case_table_names=%v\nlog_timestamps = 1\nskip-external-locking\n#skip_name_resolve = 1\n#skip-networking = 1\ndefault-storage-engine = InnoDB\n\n#timeout\nwait_timeout=28800\nlock_wait_timeout=3600\ninteractive_timeout=28800\nconnect_timeout = 20\nserver-id       = %v\n#plugin\n#plugin-load=\"semisync_master.so;semisync_slave.so\"\n\n#########SSL#############\n#ssl-ca = /home/storage/mysql_3306/data/ca.pem\n#ssl-cert = /home/storage/mysql_3306/data/server-cert.pem\n#ssl-key = /home/storage/mysql_3306/data/server-key.pem\n\n#########undo#############\ninnodb_undo_logs  =126\ninnodb_undo_directory =%v/logs\ninnodb_max_undo_log_size = 1G\ninnodb_undo_tablespaces = 8\ninnodb_undo_log_truncate = 1\ninnodb_purge_rseg_truncate_frequency = 128\n\n#########error log#############\nlog-error = %v/logs/error.log\nlog_error_verbosity  = 3\n\n#########general log#############\ngeneral_log_file=%v/logs/general.log\n\n#########slow log#############\nslow_query_log = 1\nlong_query_time=1\nslow_query_log_file = %v/logs/mysql.slow\n\n############# for replication###################\nlog-bin     = %v/logs/mysql-bin\nbinlog_format = ROW\nmax_binlog_size = 1024M\nbinlog_cache_size = 5M\nmax_binlog_cache_size = 5000M\nexpire_logs_days = 15\nslave-net-timeout=30\nlog-slow-slave-statements =1\nlog_bin_trust_function_creators = 1\nlog-slave-updates = 1\nread_only=%v\n#skip-slave-start = 1\n#super_read_only =1\n\n#relay log\nrelay-log = %v/logs/mysql-relay\nrelay-log-index=%v/logs/relay-bin.index\nmax-relay-log-size = 1024M\nrelay_log_purge = 1\n\nsync_master_info = 1\nsync_relay_log_info = 1\nsync_relay_log = 1\nrelay_log_recovery = 1\n\n#semisync\n#rpl_semi_sync_master_enabled = 1\n#rpl_semi_sync_master_wait_no_slave = 1\n#rpl_semi_sync_master_timeout = 1000\n#rpl_semi_sync_slave_enabled = 1\n#rpl_semi_sync_master_timeout = 100000000\n#rpl_semi_sync_master_wait_point = 'after_sync'\n#rpl_semi_sync_master_wait_for_slave_count = 1\n\n#ignore\n#replicate-ignore-db = 'db,'db1'\n#replicate-do-db = 'db','db1'\n#replicate-do-table = 'db.t'\n#replicate-ignore-table= 'db.t'\n\n#Multi-threaded Slave\nslave_parallel_workers=8\nslave-parallel-type=LOGICAL_CLOCK\nmaster_info_repository=TABLE\nrelay_log_info_repository=TABLE\nslave_pending_jobs_size_max=200000000\n#binlog_group_commit_sync_delay=1000\n#binlog_group_commit_sync_no_delay_count =100\n#slave_preserve_commit_order=1\n# GTID setting\ngtid-mode                      = %v\nenforce-gtid-consistency       = true\nsync-master-info               = 1\nslave-parallel-workers         = 8\nbinlog-checksum                = CRC32\nmaster-verify-checksum         = 1\nslave-sql-verify-checksum      = 1\nbinlog-rows-query-log_events   = 1\n#slave-skip-errors=1007,1051,1062\n\n#######per_thread_buffers#####################\nmax_connections=5000\nmax_user_connections=3000\nmax_connect_errors=1000000\n#myisam_recover\nmax_allowed_packet = 128M\ntable_open_cache = 6144\ntable_definition_cache = 6144\ntable_open_cache_instances = 64\n\nread_buffer_size = 1M\njoin_buffer_size = 4M\nread_rnd_buffer_size = 1M\n\n#myisam\nsort_buffer_size = 128K\nmyisam_max_sort_file_size = 10G\n#myisam_repair_threads = 1\nkey_buffer_size = 64M\nmyisam_sort_buffer_size = 32M\ntmp_table_size = 64M\nmax_heap_table_size = 64M\nquery_cache_type=0\nquery_cache_size = 0\nbulk_insert_buffer_size = 32M\nthread_cache_size = 64\n#thread_concurrency = 32\nthread_stack = 192K\n\n###############InnoDB###########################\ninnodb_data_home_dir = %v/data\ninnodb_log_group_home_dir = %v/logs\ninnodb_data_file_path = ibdata1:1000M:autoextend\ninnodb_temp_data_file_path = ibtmp1:12M:autoextend\ninnodb_buffer_pool_size = %v\ninnodb_buffer_pool_instances    = 8\ninnodb_log_file_size = 120M\ninnodb_log_buffer_size = 16M\ninnodb_log_files_in_group = 3\ninnodb_flush_log_at_trx_commit = 2\nsync_binlog = 1\ninnodb_lock_wait_timeout = 10\ninnodb_sync_spin_loops = 40\ninnodb_max_dirty_pages_pct = 80\ninnodb_support_xa = 1\ninnodb_thread_concurrency = 0\ninnodb_thread_sleep_delay = 500\ninnodb_concurrency_tickets = 1000\ninnodb_flush_method = O_DIRECT\ninnodb_file_per_table = 1\ninnodb_read_io_threads = 16\ninnodb_write_io_threads = 16\ninnodb_io_capacity = 1000\ninnodb_flush_neighbors = 1\ninnodb_purge_threads=2\ninnodb_purge_batch_size = 32\ninnodb_old_blocks_pct=75\ninnodb_change_buffering=all\ninnodb_stats_on_metadata=OFF\ninnodb_print_all_deadlocks = 1\nperformance_schema=1\ntransaction_isolation = REPEATABLE-READ\n#innodb_force_recovery=0\n#innodb_fast_shutdown=1\n#innodb_status_output=1\n#innodb_status_output_locks=1\n#innodb_status_file = 1\nsql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION\n\nreport_host='%v'\nreport_port=%v\nskip_name_resolve=on\n[mysqldump]\nquick\nmax_allowed_packet = 128M\n\n[mysql]\nno-auto-rehash\nmax_allowed_packet = 128M\nprompt                         = '\\u@\\h:\\p [\\d]> '\ndefault_character_set          = utf8\n\n[myisamchk]\nkey_buffer_size = 64M\nsort_buffer_size = 512k\nread_buffer = 2M\nwrite_buffer = 2M\n\n[mysqlhotcopy]\ninteractive-timeout\n\n[mysqld_safe]\n#malloc-lib= /usr/lib/libjemalloc.so",
		port,dataDirFull,port,baseDir,dataDirFull,dataDirFull,dataDirFull,dataDirFull,lowerCaseTableNames,serverId,dataDirFull,dataDirFull,dataDirFull,dataDirFull,dataDirFull,readOnly,dataDirFull,dataDirFull,gtidMode,dataDirFull,dataDirFull,innodbBufferPoolSize,remoteMasterAddr,port)
	fmt.Println(mysqlConfig)

	mysqlconfigFileName := fmt.Sprintf("mysql_%v.cnf",port)
	file, err := os.OpenFile(mysqlconfigFileName,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666)
	if err != nil {
		log.Printf("打开MySQL配置文件%v失败：%s",mysqlconfigFileName,err)
		return err,mysqlconfigFileName
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(mysqlConfig)

	if err != nil {
		fmt.Println("写入文件失败：", err)
		return err,mysqlconfigFileName
	}

	// 刷新缓存
	err = writer.Flush()
	if err != nil {
		log.Println("刷新缓存失败：", err)
		return err,mysqlconfigFileName
	}

	log.Println("MySQL配置文件写入成功！")
	return nil,mysqlconfigFileName
}


func initMysql (baseDir, mysqlConfigPath string,client *ssh.Client) error {
	// /usr/local/mysql/bin/mysqld --defaults-file=/home/storage/mysql_3306/mysql_3306.cnf --initialize-insecure --user=mysql
	//mysqldPath := fmt.Sprintf("%s/bin/mysqld",baseDir)
	mysqldPath := filepath.Join(baseDir,"bin/mysqld")
	defaultsFile := fmt.Sprintf("--defaults-file=%s",mysqlConfigPath)

	log.Printf("mysqldPath :%s\n", mysqldPath)
	log.Printf("configFilePath :%s\n", defaultsFile)
	cmd := fmt.Sprintf("%s %s --initialize-insecure --user=mysql",mysqldPath,defaultsFile)
	log.Println(cmd)
	if err,output:= execCmd(client,cmd);err != nil {
		log.Println("mysql初始化失败")
		log.Println(output)
		return err
	}

	log.Println("mysql初始化成功")
	return nil
}

func startMysql (baseDir, mysqlConfigPath string,client *ssh.Client) error {
	// /usr/local/mysql/bin/mysqld_safe --defaults-file=/data/storage/mysql_3306/mysql_3306.cnf --user=mysql &
	//mysqldSafePath := fmt.Sprintf("%s/bin/mysqld_safe", baseDir)

	mysqldSafePath := filepath.Join(baseDir,"bin/mysqld_safe")
	defaultsFile := fmt.Sprintf("--defaults-file=%s",mysqlConfigPath)
	cmd := fmt.Sprintf("%s %s --user=mysql >  /dev/null 2>&1 &",mysqldSafePath,defaultsFile)
	fmt.Println(cmd)
	err,output:= execCmd(client,cmd)
	log.Println(output)
	if err != nil {
		log.Println("mysql初始化失败")
		return err
	}

	log.Println("mysql启动成功")
	return nil

}

func installMysql(remoteAddr string){
	cmd := "echo 'SSH is Success!'"

	// 测试 SSH 通不通
	err,client:= createSshClient(remoteAddr)
	if err != nil {
		log.Println("创建clent失败")
		return
	}

	defer client.Close()

	// 检查数据库服务器存放安装包的目录是否存在
	remoteFilePath := filepath.Dir(remoteFile)
	fmt.Println(remoteFilePath)

	cmd = fmt.Sprintf("[ -d %s ] && echo 1 || echo 0", remoteFilePath)

	fmt.Println(cmd)

	err, output := execCmd(client,cmd)
	if err != nil {
		log.Println("执行检查目录的命令失败:", err)
	}

	if strings.TrimRight(output,"\n") == "0" {
		fmt.Println("目录不存在")
		err, output := execCmd(client,fmt.Sprintf("mkdir -p %s",remoteFilePath))

		if err!= nil {
			fmt.Println("创建存放mysql安装包目录失败,",err)
			return
		}
		fmt.Println(output)
		fmt.Println("目标端存放安装包目录已创建")
	}



		// 从中控机拷贝安装包到数据库服务器
		err = copyFile("root",mysqlPackage,remoteFile,remoteAddr)
		if err != nil {
			log.Printf("文件传输失败，%s",err)
			return
		}

		log.Printf("传输文件%s成功",mysqlPackage)


}


// CHANGE MASTER TO
//  MASTER_HOST='10.90.48.31',
//  MASTER_USER='repl_delay',
//  MASTER_PASSWORD='X0LiTey0ReEYkkN8',
//  MASTER_PORT=13306,
//MASTER_AUTO_POSITION=1;

func createRepl(slaveAddr ,masterAddr string)  {

	dataDirFullPath := filepath.Join(dataDir,fmt.Sprintf("mysql_%d",port))
	err,slaveClient:= createSshClient(slaveAddr)
	if err != nil {
		log.Println("创建clent失败")
		return
	}

	defer slaveClient.Close()
	if gtidMode=="ON" {
		cmd := fmt.Sprintf("%s/bin/mysql --socket=%s/run/mysql.sock --port=%d  -uroot -e \"CHANGE MASTER TO MASTER_HOST='%s',MASTER_PORT=%d, MASTER_USER='repl',MASTER_PASSWORD='repl',MASTER_AUTO_POSITION=1;\"",baseDir,dataDirFullPath,port,masterAddr,port)
		fmt.Println(cmd)
		log.Println("主从建立复制")
		execCmd(slaveClient,cmd)

		cmd = fmt.Sprintf("%s/bin/mysql --socket=%s/run/mysql.sock --port=%d  -uroot -e \"start slave;\"",baseDir,dataDirFullPath,port)
		fmt.Println(cmd)
		execCmd(slaveClient,cmd)
	}


}

func testMysqlConn(remoteAddr string)  {
	dataDirFullPath := filepath.Join(dataDir,fmt.Sprintf("mysql_%d",port))
	err,client:= createSshClient(remoteAddr)
	if err != nil {
		log.Printf("创建clent失败:%s",err)
		return
	}
	defer client.Close()
	cmd := fmt.Sprintf("%s/bin/mysql --socket=%s/run/mysql.sock --port=%d  -uroot -e \"select 1\"",baseDir,dataDirFullPath,port)
	for i:= 1; i<=10 ; i++ {
		if err,_ := execCmd(client,cmd);err == nil {
			fmt.Printf("启动后第%d连接MySQL成功")
			break
		} else {
			fmt.Printf("执行select 1 报错 ： %s",err)
			time.Sleep(1*time.Second)
		}
	}


}