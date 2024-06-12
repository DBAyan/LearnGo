package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
)

var (
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
)

func main()  {
	flag.BoolVar(&version,"version",false,"安装MySQL脚本版本")
	flag.BoolVar(&help,"help",false,"帮助信息")
	flag.IntVar(&port,"port",3306,"mysql端口")
	flag.StringVar(&innodbBufferPoolSize,"innodb_buffer_pool_size","1G","innodb_buffer_pool_size")
	flag.IntVar(&serverId,"server-id",0,"mysql server-id")
	flag.StringVar(&baseDir,"basedir","/usr/local/mysql","mysql安装目录")
	flag.StringVar(&dataDir,"datadir","/home/storage/","mysql数据目录")
	flag.IntVar(&readOnly,"read_only",0,"是否只读,1-只读，0-读写")
	flag.IntVar(&lowerCaseTableNames,"lower_case_table_names",0,"表名列名大小写敏感，0-敏感，1-不敏感")
	flag.StringVar(&gtidMode,"gtid-mode","OFF","GTID模式，ON-开启，OFF-关闭")
	flag.BoolVar(&onlyInit,"only-init",false,"添加该参数后不会再安装mysql软件，只会进行初始化")
	flag.Parse()
	if version {
		fmt.Println("Install MySQL Go Script v1.0")
		return
	}
	if help {
		flag.PrintDefaults()
		return

	}

	if checkPort(port) {
		log.Fatalln(port,"端口被占用")
	}
	log.Println("port:",port)
	// 如果没有设置server-id
	if serverId==0 {
		err, ipAddr, ipSuffix := getIpAddr()
		if  err != nil{
			fmt.Println("获取本机IP地址失败:",err)
		}
		log.Printf("本机IP地址：%v\n", ipAddr)
		log.Printf("本机IP地址后两位：%v\n", ipSuffix)
		serverIdStr := ipSuffix + strconv.Itoa(port)
		//log.Println("server-id:",serverIdStr)
		serverId,_ = strconv.Atoi(serverIdStr)
		log.Println("server-id:",serverId)
	}

	dataDirPort := fmt.Sprintf("mysql_%d",port)
	if baseDir == "" {
		log.Fatalln("basedir必须配置")
	} else {
		if onlyInit {
			log.Println("只做初始化，不再检查basedir目录是否为空")
		} else {
			if err := CheckDirectoryExists(baseDir); err != nil {
				log.Fatalf("检测basedir失败")
			}
		}

	}
	if dataDir == "" {
		log.Fatalln("datadir必须配置")
	} else {
		dataDir = filepath.Join(dataDir,dataDirPort)
		if !checkDir(dataDir) {
			log.Fatalf("检测datadir失败")
		}
	}

	checkAndCreateGroup("mysql")
	checkAndCreateUser("mysql")

	mysqlConfig := fmt.Sprintf(
		"[client]\nport            = %v\nsocket            = %v/run/mysql.sock\n\n# The MySQL server\n[mysqld]\n#########Basic##################\nexplicit_defaults_for_timestamp=true\nport            = %v\nuser            = mysql\nbasedir         = %v\ndatadir         = %v/data\ntmpdir          = %v/tmp\npid-file        = %v/run/mysql.pid\nsocket            = %v/run/mysql.sock\n#skip-grant-tables\n#character set\ncharacter_set_server = utf8mb4\nopen_files_limit = 65535\nback_log = 500\n#event_scheduler = ON\nlower_case_table_names=%v\nlog_timestamps = 1\nskip-external-locking\n#skip_name_resolve = 1\n#skip-networking = 1\ndefault-storage-engine = InnoDB\n\n#timeout\nwait_timeout=28800\nlock_wait_timeout=3600\ninteractive_timeout=28800\nconnect_timeout = 20\nserver-id       = %v\n#plugin\n#plugin-load=\"semisync_master.so;semisync_slave.so\"\n\n#########SSL#############\n#ssl-ca = /home/storage/mysql_3306/data/ca.pem\n#ssl-cert = /home/storage/mysql_3306/data/server-cert.pem\n#ssl-key = /home/storage/mysql_3306/data/server-key.pem\n\n#########undo#############\ninnodb_undo_logs  =126\ninnodb_undo_directory =%v/logs/\ninnodb_max_undo_log_size = 1G\ninnodb_undo_tablespaces = 8\ninnodb_undo_log_truncate = 1\ninnodb_purge_rseg_truncate_frequency = 128\n\n#########error log#############\nlog-error = %v/logs/error.log\nlog_error_verbosity  = 3\n\n#########general log#############\ngeneral_log_file=%v/logs/general.log\n\n#########slow log#############\nslow_query_log = 1\nlong_query_time=1\nslow_query_log_file = %v/logs/mysql.slow\n\n############# for replication###################\nlog-bin     = %v/logs/mysql-bin\nbinlog_format = ROW\nmax_binlog_size = 1024M\nbinlog_cache_size = 5M\nmax_binlog_cache_size = 5000M\nexpire-logs-days = 15\nslave-net-timeout=30\nlog-slow-slave-statements =1\nlog_bin_trust_function_creators = 1\nlog-slave-updates = 1\nread_only=%v\n#skip-slave-start = 1\n#super_read_only =1\n\n#relay log\nrelay-log = %v/logs/mysql-relay\nrelay-log-index=%v/logs/relay-bin.index\nmax-relay-log-size = 1024M\nrelay_log_purge = 1\n\nsync_master_info = 1\nsync_relay_log_info = 1\nsync_relay_log = 1\nrelay_log_recovery = 1\n\n#semisync\n#rpl_semi_sync_master_enabled = 1\n#rpl_semi_sync_master_wait_no_slave = 1\n#rpl_semi_sync_master_timeout = 1000\n#rpl_semi_sync_slave_enabled = 1\n#rpl_semi_sync_master_timeout = 100000000\n#rpl_semi_sync_master_wait_point = 'after_sync'\n#rpl_semi_sync_master_wait_for_slave_count = 2\n\n#ignore\n#replicate-ignore-db = 'db,'db1'\n#replicate-do-db = 'db','db1'\n#replicate-do-table = 'db.t'\n#replicate-ignore-table= 'db.t'\n\n#Multi-threaded Slave\nslave_parallel_workers=8\nslave-parallel-type=LOGICAL_CLOCK\nmaster_info_repository=TABLE\nrelay_log_info_repository=TABLE\nslave_pending_jobs_size_max=200000000\n#binlog_group_commit_sync_delay=1000\n#binlog_group_commit_sync_no_delay_count =100\n#slave_preserve_commit_order=1\n# GTID setting\ngtid-mode                      = %v\nenforce-gtid-consistency       = true\nsync-master-info               = 1\nslave-parallel-workers         = 8\nbinlog-checksum                = CRC32\nmaster-verify-checksum         = 1\nslave-sql-verify-checksum      = 1\nbinlog-rows-query-log_events   = 1\n#slave-skip-errors=1007,1051,1062\n\n#######per_thread_buffers#####################\nmax_connections=3000\nmax_user_connections=2000\nmax_connect_errors=1000000\n#myisam_recover\nmax_allowed_packet = 128M\ntable_open_cache = 6144\ntable_definition_cache = 6144\ntable_open_cache_instances = 64\n\nread_buffer_size = 1M\njoin_buffer_size = 4M\nread_rnd_buffer_size = 1M\n\n#myisam\nsort_buffer_size = 128K\nmyisam_max_sort_file_size = 10G\nmyisam_repair_threads = 1\nkey_buffer_size = 64M\nmyisam_sort_buffer_size = 32M\ntmp_table_size = 64M\nmax_heap_table_size = 64M\nquery_cache_type=0\nquery_cache_size = 0\nbulk_insert_buffer_size = 32M\nthread_cache_size = 64\n#thread_concurrency = 32\nthread_stack = 192K\n\n###############InnoDB###########################\ninnodb_data_home_dir = %v/data\ninnodb_log_group_home_dir = %v/logs\ninnodb_data_file_path = ibdata1:1000M:autoextend\ninnodb_temp_data_file_path = ibtmp1:12M:autoextend\ninnodb_buffer_pool_size = %v\ninnodb_buffer_pool_instances    = 8\ninnodb_log_file_size = 120M\ninnodb_log_buffer_size = 16M\ninnodb_log_files_in_group = 3\ninnodb_flush_log_at_trx_commit = 2\nsync_binlog = 1\ninnodb_lock_wait_timeout = 10\ninnodb_sync_spin_loops = 40\ninnodb_max_dirty_pages_pct = 80\ninnodb_support_xa = 1\ninnodb_thread_concurrency = 0\ninnodb_thread_sleep_delay = 500\ninnodb_concurrency_tickets = 1000\ninnodb_flush_method = O_DIRECT\ninnodb_file_per_table = 1\ninnodb_read_io_threads = 16\ninnodb_write_io_threads = 16\ninnodb_io_capacity = 1000\ninnodb_flush_neighbors = 1\ninnodb_purge_threads=2\ninnodb_purge_batch_size = 32\ninnodb_old_blocks_pct=75\ninnodb_change_buffering=all\ninnodb_stats_on_metadata=OFF\ninnodb_print_all_deadlocks = 1\nperformance_schema=1\ntransaction_isolation = READ-COMMITTED\n#innodb_force_recovery=0\n#innodb_fast_shutdown=1\n#innodb_status_output=1\n#innodb_status_output_locks=1\n#innodb_status_file = 1\nsql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION\n\n[mysqldump]\nquick\nmax_allowed_packet = 128M\n\n[mysql]\nno-auto-rehash\nmax_allowed_packet = 128M\nprompt                         = '\\u@\\h:\\p [\\d]> '\ndefault_character_set          = utf8\n\n[myisamchk]\nkey_buffer_size = 64M\nsort_buffer_size = 512k\nread_buffer = 2M\nwrite_buffer = 2M\n\n[mysqlhotcopy]\ninteractive-timeout\n\n[mysqld_safe]\n#malloc-lib= /usr/lib/libjemalloc.so",
		port,dataDir,port,baseDir,dataDir,dataDir,dataDir,dataDir,lowerCaseTableNames,serverId,dataDir,dataDir,dataDir,dataDir,dataDir,readOnly,dataDir,dataDir,gtidMode,dataDir,dataDir,innodbBufferPoolSize)
	fmt.Println(mysqlConfig)
	subDirs := []string{"data","tmp","logs", "run"}
	for _,subDir := range subDirs {
		path := filepath.Join(dataDir,subDir)
		err := os.MkdirAll(path,os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory %s: %s\n", path, err)
		}
		log.Printf("创建目录%s成功",path)
	}

	mysqlconfigFileName := fmt.Sprintf("mysql_%v.cnf",port)
	mysqlConfigPath := path.Join(dataDir,mysqlconfigFileName)
	fmt.Println(mysqlConfigPath)

	file, err := os.OpenFile(mysqlConfigPath,os.O_WRONLY|os.O_CREATE|os.O_EXCL,0666)
	if err != nil {
		log.Fatalf("打开MySQL配置文件%v失败",mysqlConfigPath)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(mysqlConfig)

	if err != nil {
		fmt.Println("写入文件失败：", err)
		return
	}

	// 刷新缓存
	err = writer.Flush()
	if err != nil {
		log.Println("刷新缓存失败：", err)
		return
	}

	log.Println("MySQL配置文件写入成功！")

	if onlyInit {
		log.Println("不再安装MySQL软件")
	} else {
		cmd := exec.Command("tar", "-zxvf", "mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		} else {
			log.Println("安装包解压成功！")
		}
		cmd = exec.Command("mv", "mysql-5.7.28-linux-glibc2.12-x86_64", baseDir)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		} else {
			log.Println("安装目录移动成功！")
		}
	}

	chownPath(baseDir)
	chownPath(dataDir)

	// /usr/local/mysql/bin/mysqld --defaults-file=/home/storage/mysql_3306/mysql_3306.cnf --initialize-insecure --user=mysql
	//mysqldPath := fmt.Sprintf("%s/bin/mysqld",baseDir)
	mysqldPath := filepath.Join(baseDir,"bin/mysqld")

	defaultsFile := fmt.Sprintf("--defaults-file=%s",mysqlConfigPath)

	log.Printf("mysqldPath :%s", mysqldPath)
	log.Printf("configFilePath :%s", defaultsFile)

	cmd := exec.Command(mysqldPath,defaultsFile, "--initialize-insecure" ,"--user=mysql")
	err = cmd.Run()
	if err != nil {
		log.Fatalln("mysql初始化失败")
	}

	log.Println("mysql初始化成功")


	// /usr/local/mysql/bin/mysqld_safe --defaults-file=/data/storage/mysql_3306/mysql_3306.cnf --user=mysql &
	//mysqldSafePath := fmt.Sprintf("%s/bin/mysqld_safe", baseDir)
	mysqldSafePath := filepath.Join(baseDir,"bin/mysqld_safe")
	log.Printf("mysqldSafePath:%s", mysqldSafePath)

	cmd = exec.Command(mysqldSafePath,defaultsFile,"--user=mysql")
	// 后台执行
	err = cmd.Start()
	if err != nil {
		log.Fatalln("mysql启动失败")
	}

	log.Println("mysql启动成功")

}

func getIpAddr() (err error,ipAddr string,ipSuffix string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		//fmt.Println("获取本机IP地址失败:", err)
		return err,"",""
	}

	//fmt.Println(addrs)
	for _, addr := range addrs {
		//fmt.Println(addr)
		if  ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				//fmt.Println("本机IP地址：", ipNet.IP.String())
				ipAddr = ipNet.IP.String()
				ip := net.ParseIP(ipAddr)
				ipv4 := ip.To4()
				lastByte := ipv4[3]
				secondLastByte := ipv4[2]
				ipSuffix = fmt.Sprintf("%d%d",secondLastByte,lastByte)
				// TODO 通过位运算获取IP地址 研究下
				//ipSuffix = int(secondLastByte)<<8|int(lastByte)
				break
			}
		}
	}
	if ipAddr == "" {
		return fmt.Errorf("未找到有效的本机 IP 地址"), "",""
	}
	return nil,ipAddr,ipSuffix
}


func checkPort(port int) bool {
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{IP:net.IPv4(127,0,0,1),Port:port})
	if err != nil {
		log.Printf("检测端口%d监听失败：%v", port,err)
		return true
	}
	defer ln.Close()
	log.Printf("检测端口%d监听成功",port)
	return false
}

func CheckDirectoryExists(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		return fmt.Errorf("Directory already exists")
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func checkDir(path string) bool {
	_,err := os.Stat(path)
	if err!= nil{
		if os.IsNotExist(err) {
			log.Printf("检测目录%v不存在",path)
			return true
		} else {
			log.Printf("检测目录%v报错%v",path,err)
			return false
		}
	} else {
		file, err := os.ReadDir(path)
		if err !=nil {
			log.Printf("读取目录%v失败：%s",path,err)
		}
		if len(file) == 0 {
			log.Printf("检测目录%v为空",path)
			return true
		} else {
			log.Printf("检测目录%v不为空",path)
			return false
		}
	}
}

func checkAndCreateUser(userName string)  {
	_,err := user.Lookup(userName)
	if err != nil  {
		fmt.Printf("用户%v不存在\n", userName)

		cmd := exec.Command("useradd","-r","-g",userName,"-s", "/bin/false", "mysql")
		if err:= cmd.Run(); err != nil {
			log.Fatalf("创建用户 %s失败:%s", userName,err)
		}
		log.Printf("创建用户 %s成功", userName)
	} else {
		fmt.Printf("用户%v已存在\n", userName)
	}

}

//useradd -r -g mysql -s /bin/false mysql
func checkAndCreateGroup(groupName string)   {
	_,err := user.LookupGroup(groupName)
	if err != nil {
		log.Printf("用户组%v不存在\n",groupName)

		cmd := exec.Command("groupadd","mysql")
		if err := cmd.Run(); err!=nil {
			log.Fatalf("创建用户组%s失败",groupName)
		}
		log.Printf("创建用户组%s成功",groupName)
	} else {
		log.Printf("用户组%v已存在\n",groupName)
	}
}


func installYumPkg(packageName string){
	cmd := exec.Command("yum", "install", "-y", packageName)
	if err := cmd.Run(); err!=nil {
		log.Fatalf("依赖包%v安装失败:%v",packageName,err)
	} else {
		log.Printf("依赖包%v 安装成功", packageName)
	}
}
func chownPath(path string)  {
	cmd := exec.Command("chown","-R","mysql:mysql", path)
	if err := cmd.Run(); err!=nil {
		log.Fatalf("修改目录%v用户和组为mysql:mysql失败%v",path,err)
	} else {
		log.Printf("修改目录%v用户和组为mysql:mysql成功", path)
	}
}