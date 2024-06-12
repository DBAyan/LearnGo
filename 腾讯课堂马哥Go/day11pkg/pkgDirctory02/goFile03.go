package pkgDirctory02

import (
	"fmt"

	mypkg "day11pkg/pkgDirctory01" // 别名 module_name/目录名称
)

type User struct {
	Name string
	Age int
	Address mypkg.Address
}

func init()  {
	fmt.Println("in package pkgD02 bbb ")
}

