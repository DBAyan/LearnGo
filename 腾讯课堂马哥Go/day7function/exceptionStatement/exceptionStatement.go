package main

import (
	"errors"
	"fmt"
	"time"
)

type PathError struct {
	path string
	op string
	createtime time.Time
	message string
}

//func NewPathErrors(path,op, message string) PathError {
//
//}

func divide(a,b int) (int, error) {
	if b == 0 {
		return -1, errors.New("divide by zero")
	}
	return a / b, nil
}

func main()  {

	res , err := divide(2,1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}