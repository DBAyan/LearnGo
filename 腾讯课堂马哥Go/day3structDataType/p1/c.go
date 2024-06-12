package p1

import "fmt"

var userinfo UserInfo

func userfunc()  float64 {
	userinfo.Score = 100
	fmt.Println(userinfo.Score)

	sc := shopcar{id: 1, money: 15.6}
	fmt.Println(sc)
	return userinfo.Score
}


