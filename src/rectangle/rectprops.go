package rectangle

import (
	"fmt"
	"math"
)

func init(){
	fmt.Println("rectangle package initialized")
}

// 可以被导出的函数的首字母必须大写

func Area(len, wid float64) float64{
	area := len * wid
	return area
}

func Diagonal(len,wid float64) float64{
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
