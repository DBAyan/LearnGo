package main


//.：匹配任意字符，除了换行符。
//*：匹配前面的字符零次或多次。
//+：匹配前面的字符一次或多次。
//?：匹配前面的字符零次或一次。
//[]：匹配方括号内的任意一个字符。
//[^]：匹配除了方括号内的任意一个字符。
//()：定义一个捕获组，可以使用 | 连接多个捕获组。
//\d：匹配数字字符，等同于 [0-9]。
//\s：匹配空白字符，包括空格、制表符和换行符等。
//\w：匹配单词字符，包括字母、数字和下划线等。
//^：匹配输入字符串的开头。
//$：匹配输入字符串的结尾。
//\b：匹配单词边界。
//  (?i) 大小写不敏感
// \s+ 一个或多个空格
// \S 表示匹配任意非空白字符
import (
	"fmt"
	"regexp"
	"strings"
)

var (
	alterTableExplicitSchemaTableRegexps = []*regexp.Regexp{
		// ALTER TABLE `scm`.`tbl` something
		// (?i)\balter\s+table\s+`([^`]+)`[.]`([^`]+)`\s+(.*$)
		regexp.MustCompile(`(?i)\balter\s+table\s+` + "`" + `([^` + "`" + `]+)` + "`" + `[.]` + "`" + `([^` + "`" + `]+)` + "`" + `\s+(.*$)`),
		// ALTER TABLE `scm`.tbl something
		regexp.MustCompile(`(?i)\balter\s+table\s+` + "`" + `([^` + "`" + `]+)` + "`" + `[.]([\S]+)\s+(.*$)`),
		// ALTER TABLE scm.`tbl` something
		regexp.MustCompile(`(?i)\balter\s+table\s+([\S]+)[.]` + "`" + `([^` + "`" + `]+)` + "`" + `\s+(.*$)`),
		// ALTER TABLE scm.tbl something
		regexp.MustCompile(`(?i)\balter\s+table\s+([\S]+)[.]([\S]+)\s+(.*$)`),
	}
)

// 3个捕获组
var r1 = regexp.MustCompile(`(?i)\balter\s+table\s+([\S]+)[.]([\S]+)\s+(.*$)`)

func main()  {
	r := regexp.MustCompile(`\d+`)
	res := r.FindString("价格是2.600")
	fmt.Println(res)
	fmt.Println(`(?i)\balter\s+table\s+` + "`" + `([^` + "`" + `]+)` + "`" + `[.]` + "`" + `([^` + "`" + `]+)` + "`" + `\s+(.*$)`)

	//func (re *Regexp) FindStringSubmatch(s string) []string
	//其中，re 是一个编译后的正则表达式对象，s 是要匹配的字符串。函数返回一个字符串切片，其中第一个元素是与整个正则表达式匹配的子字符串，后面的元素则是每个捕获组匹配的结果。
	strbyte := r1.FindStringSubmatch("alter table scm_name.tbl_name add index idx_a(a);")
	fmt.Println(strbyte)
	fmt.Println(len(strbyte))
	for idx,val := range strbyte {
		fmt.Printf("%v:%v\n",idx, val)
	}
	// 返回
	//0:alter table scm_name.tbl_name add index idx_a(a); 第一个元素是与整个正则表达式匹配的子字符串
	//1:scm_name  后面的元素则是每个捕获组匹配的结果
	//2:tbl_name
	//3:add index idx_a(a);

	tokens := tokenizeAlterStatement("alter table scm_name.tbl_name add index idx_a(a);")
	fmt.Println(tokens)
	fmt.Println(len(tokens))

}


func  tokenizeAlterStatement(alterStatement string) (tokens []string) {
	terminatingQuote := rune(0)
	fmt.Println(terminatingQuote)
	f := func(c rune) bool {
		switch {
		case c == terminatingQuote:
			terminatingQuote = rune(0)
			return false
		case terminatingQuote != rune(0):
			return false
		case c == '\'':
			terminatingQuote = c
			return false
		case c == '(':
			terminatingQuote = ')'
			return false
		default:
			return c == ','
		}
	}

	tokens = strings.FieldsFunc(alterStatement, f)
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	return tokens
}