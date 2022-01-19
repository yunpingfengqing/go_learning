package controller

import (
	"blog/src/bean"
	"blog/src/service"
	"blog/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InsertArticle(c *gin.Context) {
	var article bean.Article
	err := c.Bind(&article)

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		err = service.InsertArticle(&article)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, "插入数据成功")
		}
	}
}

func DeleteArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		err = service.DeleteArticleById(id)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, "删除数据成功")
		}
	}
}

func UpdateArticle(c *gin.Context) {
	var article bean.Article
	err := c.BindJSON(&article)

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		err = service.UpdateArticle(&article)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, "更新数据成功")
		}
	}
}

func QueryArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		var article bean.Article
		article, err = service.QueryArticleById(id)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, article)
		}
	}
}

func QueryHotArticleList(c *gin.Context) {
	var pageSize, pageNum, typeId int
	var err error

	pageSize, err = strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	pageNum, err = strconv.Atoi(c.Query("pageNum"))
	if err != nil || pageNum < 1 {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	typeId, err = strconv.Atoi(c.Query("typeId"))
	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var articleList []bean.Hot_Article
	articleList, err = service.QueryHotArticleList(pageSize, pageNum, typeId)

	if err != nil {
		util.Error(c, http.StatusForbidden, err.Error())
	} else {
		util.Success(c, articleList)
	}
}

func QueryHotArticleListByRedis(c *gin.Context) {
	var pageSize, pageNum, typeId int
	var err error

	pageSize, err = strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	pageNum, err = strconv.Atoi(c.Query("pageNum"))
	if err != nil || pageNum < 1 {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	typeId, err = strconv.Atoi(c.Query("typeId"))
	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var articleList []string
	articleList, err = service.QueryHotArticleListByRedis(pageSize, pageNum, typeId)

	if err != nil {
		util.Error(c, http.StatusForbidden, err.Error())
	} else {
		util.Success(c, articleList)
	}
}