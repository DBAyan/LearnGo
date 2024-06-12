package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 字符串拼接
func arr2string(arr []int) string  {
	var sb strings.Builder

	for _, ele := range arr {
		fmt.Printf("ele is %d, type is %T\n",ele,ele)
		//sb.WriteString(strconv.Itoa(ele))
		sb.WriteString(strconv.FormatInt(int64(ele), 10))
		sb.WriteString(" ")
	}
	str := sb.String()
	return strings.Trim(str, " ")
}

func main()  {
	var s []int
	s = make([]int,0, 10)
	fmt.Printf("s:%v\n length of sli is %d\n capacity of sli is %d\n",s, len(s), cap(s))

	for i:=1;i<=100;i++ {
		s=append(s, rand.Intn(128))
	}
	fmt.Printf("s:%v\n length of sli is %d\n capacity of sli is %d\n",s, len(s), cap(s))

	var m map[int]int
	m = make(map[int]int)

	for _, ele := range s{
		//m[ele] = idx
		value, ok := m[ele]
		if ok {
			m[ele] = value + 1
		} else {
			m[ele] = 1
		}
	}
	fmt.Println(m)
	fmt.Printf("length of m is %d\n", len(m))


	res := arr2string([]int{1,2,3})
	fmt.Println(res)

}
