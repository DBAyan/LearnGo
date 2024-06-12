package main

import (
	"flag"
)
var name string

func init()  {
	flag.StringVar(&name,"name","yhh","someone")
}

func main()  {
	flag.Parse()
	hello(name)
}