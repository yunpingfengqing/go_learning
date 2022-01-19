package controller

import (
	"blog/src/bean"
	"blog/src/service"
	"blog/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	var user bean.User
	err := c.Bind(&user)

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		err = service.InsertUser(&user)
		
		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, "插入数据成功")
		}
	}
}

func DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		err = service.DeleteUserById(id)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, "删除数据成功")
		}
	}
}

func UpdateUser(c *gin.Context) {
	var user bean.User
	err := c.BindJSON(&user)

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		err = service.UpdateUser(&user)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, "更新数据成功")
		}
	}
}

func QueryUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
	} else {
		var user bean.User
		user, err = service.QueryUserById(id)

		if err != nil {
			util.Error(c, http.StatusForbidden, err.Error())
		} else {
			util.Success(c, user)
		}
	}
}
