package service

import (
	"blog/src/bean"
	"blog/src/config"
	"encoding/json"
	"strconv"
	"time"
)

func InsertArticle(article *bean.Article) (err error) {
	article.Create_Time = time.Now()
	article.Update_Time = time.Now()
	err = config.Db.Create(&article).Error
	return
}

func DeleteArticleById(id int) (err error) {
	err = config.Db.Delete(&bean.Article{}, id).Error
	return
}

func UpdateArticle(article *bean.Article) (err error) {
	article.Update_Time = time.Now()
	err = config.Db.Model(&article).Select("Article_Name", "Type_Id", "Content", "Status").Updates(article).Error
	return
}

func QueryArticleById(id int) (article bean.Article, err error) {
	err = config.Db.First(&article, "id = ?", id).Error

	// 查询文章内容的时候，往浏览表更新浏览次数
	if err == nil {
		var articleViews bean.ArticleViews
		now := time.Now()
		date := now.Format("2006-01-02")

		// 确认当天该文章是否有过浏览记录
		articleViews, err = QueryArticleViews(date, article.Id)
		// 如果该文章当天首次被浏览，插入一条新数据，否则更新浏览次数
		if err != nil {
			articleViews.Date = date
			articleViews.Article_Id = article.Id
			articleViews.Views = 1
			err = InsertArticleViews(&articleViews)
		} else {
			articleViews.Date = date
			err = UpdateArticleViews(&articleViews)
		}
	}

	return
}

func HotArticleListToString(articleList []bean.Hot_Article) (hotArticle []string, er error) {
	hotArticle = make([]string, len(articleList))
	for key, article := range articleList{
		jsons, err := json.Marshal(article)

		if err != nil {
			er = err
			return
		}

		hotArticle[key] = string(jsons)
	}
	return
}

func QueryHotArticleList(pageSize int, pageNum int, typeId int) (articleList []bean.Hot_Article, err error) {
	err = config.Db.Model(&bean.Article{}).Select("Article.Article_Name, sum(aw.Views) as Total").Joins("join Article_Views aw on aw.Article_Id = Article.Id").Where("Article.Type_Id = ?", typeId).Group("Article.Id").Order("sum(aw.Views) desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error

	return
}

func QueryHotArticleListByType(typeId int) (articleList []bean.Hot_Article, err error) {
	err = config.Db.Model(&bean.Article{}).Select("Article.Article_Name, sum(aw.Views) as Total").Joins("join Article_Views aw on aw.Article_Id = Article.Id").Where("Article.Type_Id = ?", typeId).Group("Article.Id").Order("sum(aw.Views) desc").Find(&articleList).Error

	return
}

func QueryHotArticleListByRedis(pageSize int, pageNum int, typeId int) (hotArticle []string, err error) {
	hotArticle, err = config.Redisdb.LRange("HotArticle_" + strconv.Itoa(typeId), int64((pageNum-1)*pageSize), int64(pageNum*pageSize)).Result()

	if err != nil {
		return
	}

	if len(hotArticle) == 0 {
		articleList,er := QueryHotArticleListByType(typeId)

		if er != nil {
			err = er
			return
		}

		hotArticle, err = HotArticleListToString(articleList)

		err = config.Redisdb.RPush("HotArticle_" + strconv.Itoa(typeId), hotArticle).Err()
		config.Redisdb.Incr("HotArticle_" + strconv.Itoa(typeId))
		config.Redisdb.Expire("HotArticle_" + strconv.Itoa(typeId), 1 * time.Hour)

		if err != nil {
			return
		}

		hotArticle, err = config.Redisdb.LRange("HotArticle_" + strconv.Itoa(typeId), int64((pageNum-1)*pageSize), int64(pageNum*pageSize)).Result()
	}
	return
}