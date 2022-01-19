package bean

import (
	"time"
)

type Article struct {
	Id           int `gorm:"default:uuid_generate_v3(); PRIMARY_KEY"`
	Article_Name string
	Author       int
	Type_Id      int
	Content      string
	Status       int `gorm:"default:1"`
	Create_Time  time.Time
	Update_Time  time.Time
}

type Hot_Article struct {
	Article_Name string
	Total        int
}

func (Article) TableName() string {
	return "Article"
}
