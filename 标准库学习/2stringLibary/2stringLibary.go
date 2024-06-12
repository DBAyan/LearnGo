package main

import (
	"fmt"
	"strings"
	"unicode"
)

// https://blog.csdn.net/LiuHuan303/article/details/124296264
//https://www.topgoer.cn/docs/golangstandard/golangstandard-1cmksr4dhortl#b4vnzt
/*
func Compare(a, b string) int
func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsRune(s string, r rune) bool
func Count(s, substr string) int
func EqualFold(s, t string) bool
func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Index(s, substr string) int
func IndexAny(s, chars string) int
func IndexByte(s string, c byte) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int
func Join(elems []string, sep string) string
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int
func Map(mapping func(rune) rune, s string) string
func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string
func ReplaceAll(s, old, new string) string
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string
func Title(s string) string
func ToLower(s string) string
func ToLowerSpecial(c unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(c unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(c unicode.SpecialCase, s string) string
func ToValidUTF8(s, replacement string) string
func Trim(s, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string

 */


func main()  {
	var str1 string = "hello"
	var str2 string = "world"
	str3 := "Hello"

	var res = str1 + " " + str2
	fmt.Printf("字符串：%s,长度：%d,类型：%T\n",res,len(res),res)

	// 比较字符串
	fmt.Println(strings.Compare(str1,str2))
	fmt.Println(strings.Compare(str1,str1))
	fmt.Println(strings.Compare(str2,str1))

	fmt.Println(strings.Compare(str1,str3))
	fmt.Println(strings.EqualFold(str1,str3)) // 不区分大小写

	// 字符串是否包含子串
	fmt.Println(strings.Contains(str1,"llo"))
	fmt.Println(strings.Contains(str1,"yhh"))

	// 字符串是否包含子串中任意一个Unicode码点
	fmt.Println(strings.ContainsAny(str1,"o y"))
	fmt.Println(strings.ContainsAny(str1,"abcdfdy"))

	// type rune = int32 ,unicode 码点
	fmt.Println(strings.ContainsRune(str1,104))

	// 统计字符串中子串的总数
	fmt.Println(strings.Count(str1,"l")) //2
	fmt.Println(strings.Count(str1,"")) //6
	fmt.Println(strings.Count("fivevev", "vev")) // 1


	res2 := strings.Fields("yan hai hang") // 以空格分割字符串为 []strings 切片
	fmt.Printf("split res2 is %v\n",res2)
	for _ ,val := range res2{
		fmt.Println(val)
	}

	res3 := strings.FieldsFunc("  aa   bb cc dd ",unicode.IsSpace)
	fmt.Printf("res3 is %v\n",res3)

	// 分割字符串
	res4 := strings.Split("didi|ehr|DBA","|")
	fmt.Printf("res4 is %v\n",res4)

	res5 := strings.SplitAfter("didi|ehr|DBA","|")
	fmt.Printf("%v\n", res5)

	// 前缀 或 后缀
	fmt.Println(strings.HasPrefix("Gopher","G"))
	fmt.Println(strings.HasPrefix("Gopher","g"))
	fmt.Println(strings.HasPrefix("Gopher",""))
	fmt.Println(strings.HasSuffix("Gopher","er"))
	fmt.Println(strings.HasSuffix("Gopher","e"))
	fmt.Println(strings.HasSuffix("Gopher",""))

	// 字符串中某个子串的索引位置
	fmt.Println(strings.Index("hello","l")) // 子串第一次出现的索引位置
	fmt.Println(strings.LastIndex("hello","l")) // 子串最后一次出现的索引位置

	// 字符串拼接
	name := strings.Join([]string{"yan","hai","hang"},"-")
	fmt.Println(name)

	// 字符串重复几次
	fmt.Println("ba" + strings.Repeat("na",2))

	// 大小写转换
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("WORLD"))

	// 裁剪字符串两侧的子串
	fmt.Println(strings.Trim("hello","o"))

	// Builder 结构体类型
	var b strings.Builder
	b.WriteString("a")
	fmt.Printf("%v\n",b) // {0xc0000bfed0 [97]}
	fmt.Println(b.String()) // a

}
