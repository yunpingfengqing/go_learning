package service

import (
	"blog/src/bean"
	"blog/src/config"
	"time"
)

func InsertUser(user *bean.User) (err error) {
	user.Create_Time = time.Now()
	err = config.Db.Create(&user).Error
	return
}

func DeleteUserById(id int) (err error) {
	err = config.Db.Delete(&bean.User{}, id).Error
	return 
}

func UpdateUser(user *bean.User) (err error) {
	err = config.Db.Model(&user).Select("User_Name", "Password", "Status").Updates(user).Error
	return
}

func QueryUserById(id int) (user bean.User, err error) {
	err = config.Db.First(&user, "id = ?", id).Error

	return
}
