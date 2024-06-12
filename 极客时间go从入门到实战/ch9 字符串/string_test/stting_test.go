package string_test

/*
字符串
1 字符串是数据类型，不是引用或者指针类型。 所以他的零值不是空 而是空字符串，
2 字符串是只读的byte slice（byte的切片）,len函数可以显示它所包含的byte数，与多少个字符是不一样的
3 string的bytes数组可以存放任何数据，不仅仅是可以放可见字符 比如汉字英语，
甚至于完全不属于可见字符编码的二进制数据也是可以存放的

Unicode 是一种字符集(code point）, 是一种字符的编码

UTF8是Unicode的存储形式（转换为字节序列的规则）

三 字符串的遍历
range

四 字符串的函数
字符串有很多方法 可以参考官方文档
分割 strings.split
连接 strings.Join
转换
字符串转换为整数

数字转字符串
strconv.Itoa(1000)
*/
import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T){
	var s string
	t.Log(s) //初始化为默认零值 空字符串！

	s = "hello"
	s1 := "颜"
	t.Log(s,s1)
	t.Log(len(s),len(s1)) // 5 3 byte数 英文单词1个字符占1个字节  汉字1个字符 3字节

	s = "\xE4\xB8\xA5" // 可以存放任何二进制数据
	t.Log(s)
	t.Log(len(s)) // len求出来时的byte数 ，并不一定是字符数

	s = "\xE4\xBA\xA5\xEE" //存放任意一个二进制数据
	t.Log(s)
	t.Log(len(s))

	//s[1] = "3" // 是不可变的切片 ，是不能赋值的 cannot assign to s[1] (strings are immutable)

	s2 := "中"
	c := []rune(s2) //在go语言中 rune可以取出字符串的Unicode
	t.Log(len(s2))
	t.Log(c) //  [20013]
	t.Logf("汉字'中'的unicode %x",c[0]) //  4e2d
	t.Logf("汉字'中'的UTF8 %x",s2) //  e4b8ad

	// 如果你有一个字符串 ，你想遍历每个rune的时候，也是可以使用range的。
	// "%[1]c ,%[1]d",c  [1]都是使用第一个变量对应，以%c格式化 以%d格式化的差别了
	// range与字符串配合的时候 go会默认转化输出的是rune

	s3 := "颜海航"
	c3 := []rune(s3) //  [39068 28023 33322]

	t.Log(c3)
	for _,c := range s3{
		t.Logf("%[1]c ,%[1]d, %[1]x",c) // 字符 十进制的Unicode 十六进制的Unicode
	}
// 遍历字符串时 第一个返回值为字符的下标志
	s4 := "中国人"
	for k,v := range s4{
		t.Log(k,v)
	}

}

func TestStringFunc(t *testing.T){
	s := "A,B,C"
	parts := strings.Split(s,",")
	t.Logf("%T",parts) // 字符串切片 []string
	t.Log(parts) // [A B C]
	for _,part := range parts{
		t.Log(part)
	}
	t.Log(strings.Join(parts,"-"))

}

func TestStingConv(t *testing.T){
	s := strconv.Itoa(1000)
	t.Log("str"+s)
	// t.Log(33 + strconv.Atoi("101")) // multiple-value strconv.Atoi() in single-value context


	t.Log(strconv.Atoi("10000")) // 101 <nil>

	if v,err := strconv.Atoi("10000");err == nil{
		t.Log(33 + v)
	}
}

