package service

import (
	"blog/src/bean"
	"blog/src/config"
)

func QueryTypeById(id int) (articleType bean.ArticleType, err error) {
	err = config.Db.First(&articleType, "id = ?", id).Error
	return
}
