package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {
	file,err := os.Open("/Users/didi/Desktop/源码/Learn_go/learn_self/canalMonitor/binlog")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}