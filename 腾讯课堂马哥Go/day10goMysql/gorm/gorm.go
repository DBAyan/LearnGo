package main

import (
	"time"
)


type Student struct {
	Id         int    `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name"`
	Province   string
	City       string    `gorm:"column:city"`
	Address    string    `gorm:"column:addr"`
	Score      float64   `gorm:"column:score"`
	Enrollment time.Time `gorm:"column:enrollment;type:date"`
}

func (Student) TableName() string {
	return "student"
}





