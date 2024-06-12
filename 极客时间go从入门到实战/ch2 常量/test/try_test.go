package test

/*
1 源码文件必须以_test.go结尾！！！！
2 测试方法名 必须以 Test开头 ，大写的含义还有包外可以访问 func TestXXX(t *testing.T){...}
*/
import "testing"

func TestFirstTry(t *testing.T)  {
	t.Log("my first try")
}
