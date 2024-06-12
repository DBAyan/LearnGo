package main

import (
	"fmt"
	"strings"
)

func main() {

	// 字符串赋值
	s := "中国人"
	str := "颜,aa|yan|hai,hangYy"
	s1 := "My nama is 颜海航😄"                        // 字符串里可以包含任意Unicode字符
	s2 := "He say :\"I'm fine\"\n \\Thank\tyou.\\" // 包含装转义字符
	s3 := `here is first line

there is third line` // 反引号里的转义字符无效。反引号里的原封不动的输出，好汉空白符和换行符
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// 常见字符串操作
	fmt.Printf("s1 length is %d\n", len(s1))                           // 求字符串的长度 返回的是字节的个数
	fmt.Printf("s contains string 中： %t \n", strings.Contains(s, "中")) // 包含
	fmt.Printf("s start with 中： %t\n", strings.HasPrefix(s, "中"))      //是否以某个字符串为前缀
	fmt.Printf("s end with 国 %t\n", strings.HasSuffix(s, "国"))         //是否以某个字符串为后缀
	fmt.Printf("%v\n", strings.Split(str, "|"))                        // 按照某个字符分割字符串
	fmt.Printf("%d\n", strings.Index(str, "y"))                        // 字符串出现的首个 索引的位置
	fmt.Printf("%d\n", strings.LastIndex(str, "y"))                    // 这个字符最后一次出现的额索引位置
	// 字符串拼接
	fmt.Printf("%s\n", s+str) // 使用 + 拼接
	res := fmt.Sprintf("%s", strings.Split(str, "|"))
	fmt.Printf("%v\n", res)

	res2 := strings.Join(strings.Split(str, "|"), "+") // strings.Join 拼接
	fmt.Printf("%v\n", res2)

	arr := []byte(s)                                                       // 将字符串转换为字节数组
	brr := []rune(s)                                                       // 将字符串转换为 rune(unicode)数组
	fmt.Printf("arr length is %d ,brr length is %d\n", len(arr), len(brr)) // arr length is 9 ,brr length is 3
	fmt.Printf("arr : %v, brr:%v\n", arr, brr)                             // arr : [228 184 173 229 155 189 228 186 186], brr:[20013 22269 20154] 十进制
	// 遍历字节数组
	for _, ele := range arr {
		fmt.Printf("%d ", ele)
	}
	fmt.Println("")
	// 遍历字符串数组
	for _, ele2 := range brr {
		fmt.Printf("%d ", ele2)
	}

	fmt.Println("")

	// 遍历字符串 使用 %c 格式化输出 中 国 人
	for _, ele3 := range s {
		fmt.Printf("%c ", ele3)
	}

	// 遍历字符串 使用%d返回rune  20013 22269 20154
	for _, ele4 := range s {
		fmt.Printf("%d ", ele4)
	}

	fmt.Println("")

	// int 转 byte(uint8) 超过255 会截断
	var i int = 257
	var by byte = byte(i)
	fmt.Println(by)
	fmt.Printf("type by :%T;type i :%T\n", by, i)

	// float 转 int 小数部分会丢失
	var f float64 = 3.14
	var i64 int = int(f)
	fmt.Printf("type f :%T, type i64 %T ,f:%f,i64 :%d", f, i64, f, i64)
}
