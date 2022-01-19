package bean

import (
	"time"
)

type User struct {
	Id          int `gorm:"default:uuid_generate_v3(); PRIMARY_KEY"`
	User_Name   string
	Account     string
	Password    string
	Status      int `gorm:"default:1"`
	Create_Time time.Time
}

func (User) TableName() string {
	return "User"
}
