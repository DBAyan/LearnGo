/*
pck01 comment 可以在这里写package的注释
*/

package pkgDirctory01

import "fmt"

// 地址结构体

type Address struct {
	Province string
	City string
}

func init()  {
	fmt.Println("in package pkgD01 aaa ")
}