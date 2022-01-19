package service

import (
	"blog/src/bean"
	"blog/src/config"
)

func InsertArticleViews(articleViews *bean.ArticleViews) (err error) {
	err = config.Db.Create(&articleViews).Error
	return
}

func UpdateArticleViews(articleViews *bean.ArticleViews) (err error) {
	err = config.Db.Model(&bean.ArticleViews{}).Where("Date = ? AND Article_Id = ?", articleViews.Date, articleViews.Article_Id).Update("Views", articleViews.Views+1).Error
	return
}

func QueryArticleViews(date string, articleId int) (articleViews bean.ArticleViews, err error) {
	err = config.Db.Where("Date = ? AND Article_Id = ?", date, articleId).First(&articleViews).Error
	return
}