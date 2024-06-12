package condition

import "testing"

/*
switch
1 条件表达式不限制为常量或者整数
2 单个case里，可以出现多个结果选项，使用逗号隔开
3 与C语言不同 Go语言不需要用break来明确退出一个case
4 可以不设定switch之后的条件表达式，在这种情况下,整个switch结构与多个if ... else ...的逻辑结果相同
*/
func TestSwitchMuiltCase(t *testing.T){
	for i:=0;i<5;i++{
		switch i {
		case 0,2:
			t.Logf("%d is Even",i)
		case 1,3:
			t.Logf("%d is Odd",i)
		default:
			t.Logf(" %d ,it is not 0-3",i)
		}
	}
}

func TestSwitchCaseCondition(t *testing.T){
	for i:=0;i<5;i++{
		switch  {
		case i%2==0:
			t.Logf("%d is Even",i)
		case i%2==1:
			t.Logf("%d is Even",i)
		default:
			t.Log("unknow")


		}
	}
}