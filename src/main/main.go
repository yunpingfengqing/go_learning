package main

import (
	"blog/src/bean"
	"blog/src/config"
	"blog/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	db := config.Db

	db.AutoMigrate(&bean.User{}, &bean.Article{}, &bean.ArticleType{}, &bean.ArticleViews{})

	r := gin.Default()
	// 注册路由
	Group := r.Group("/")
	{
		// 增加
		Group.POST("/user", controller.InsertUser)
		// 查看用户
		Group.GET("/user/:id", controller.QueryUserById)
		// 修改用户
		Group.PUT("/user", controller.UpdateUser)
		// 删除用户
		Group.DELETE("/user/:id", controller.DeleteUserById)

		// 增加帖子
		Group.POST("/article", controller.InsertArticle)
		// 查看帖子
		Group.GET("/article/:id", controller.QueryArticleById)
		// 查看热门帖子
		Group.GET("/articles/hot", controller.QueryHotArticleList)
		Group.GET("/articles/redisHot", controller.QueryHotArticleListByRedis)
		// 修改帖子
		Group.PUT("/article", controller.UpdateArticle)
		// 删除帖子
		Group.DELETE("/article/:id", controller.DeleteArticleById)

	}

	// 运行8080端口
	r.Run(":8080")
}
