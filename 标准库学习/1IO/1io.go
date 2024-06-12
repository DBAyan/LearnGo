package main

import (
	"fmt"
	"io"
)

// https://www.kancloud.cn/yetship/golang_standard_library_samples/527123
// http://t.zoukankan.com/wanghui-garcia-p-10314495.html

func ReadFrom(r io.Reader,num int) ([]byte,error){
	p := make([]byte,num)
	n,err := r.Read(p)
	if n > 0 {
		return p[:n],nil

	}
	return p,err
}

func main()  {
	//data,err := ReadFrom(os.Stdin,11)
	//fmt.Println(data,err)

	//d,e := ReadFrom(strings.NewReader("abcdefghijkl"), 12)
	//fmt.Println(d,e)
	b := []byte{228,184,165}
	s := string(b)
	fmt.Println(s)
}
