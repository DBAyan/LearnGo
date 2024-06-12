package main

import (
	"fmt"
	"os"
)

func main()  {
	exitWithError(fmt.Errorf("错误发生"))
}

func exitWithError(err error)  {

	fmt.Fprintf(os.Stderr,"error:%v",err)
	os.Exit(1)
}



