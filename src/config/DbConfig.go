package config

import (
	"blog/src/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	util.ProtectRun(func() {
		var err error
		dsn := "dms:learningWeb(123)@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}
	})
}
