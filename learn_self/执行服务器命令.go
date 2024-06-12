package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	cmd := exec.Command("ls","-al")
	output ,err := cmd.Output()
	if err != nil {
		fmt.Printf("执行报错！%v", err)
	} else {
		fmt.Println(string(output))
	}
}
