package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main(){
	var fptr = flag.String("FilePath", "/Users/yanhaihang/Desktop/go/LearnGo/go语言中文网/test.txt", "The path of file")
	flag.Parse()
	f,err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte,3)
	for {
		_,err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}