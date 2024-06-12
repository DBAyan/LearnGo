package main

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
)

func main()  {
	checkAndCreateGroup("testgroup")
	checkAndCreateUser("testuser")
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