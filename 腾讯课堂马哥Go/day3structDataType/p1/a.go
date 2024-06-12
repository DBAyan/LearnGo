package p1

import (
	"time"
)

type UserInfo struct {
	id int
	name,addr string // 多个字段类型相同时可以简写为一行
	Score float64
	enrollment time.Time
}


type shopcar struct {
	id int
	money float64
}