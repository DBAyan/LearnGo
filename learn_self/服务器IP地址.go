package main

import (
	"fmt"
	"net"
)

func main()  {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("获取本机IP地址失败:", err)
		return
	}

	fmt.Println(addrs)
	for _, addr := range addrs {
		fmt.Println(addr)
		if  ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("本机IP地址：", ipNet.IP.String())
			}
		}
	}
}
