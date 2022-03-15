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

func Login(user *bean.User) (err error) {
	err = config.Db.Where("Account = ? AND Password = ? AND Status = 1", user.Account, user.Password).First(&user).Error

	if err != nil {
		return
	}

	if user.Id != 0 {
		err = config.Redisdb.Set("User_"+user.Account, 1, 5*time.Minute).Err()

		if err != nil {
			return
		}

		tm, er := config.Redisdb.TTL("User_" + user.Account).Result()

		if er != nil {
			return er
		}
		panic(tm)
	}

	return
}

func Logout(account string) (err error) {
	var status string
	status, err = config.Redisdb.Get("User_" + account).Result()

	if err != nil {
		return
	}

	if status == "1" {
		err = config.Redisdb.Del("User_" + account).Err()
	}

	return
}

func DeleteUserByAccount(account string) (err error) {
	err = config.Db.Model(&bean.User{}).Where("Account = ? AND Status = 1", account).Update("Status", 0).Error
	return
}
